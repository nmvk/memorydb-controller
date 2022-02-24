from dataclasses import dataclass, field
from typing import List

import boto3
from acktest.bootstrapping import Bootstrappable
from acktest.resources import random_suffix_name


@dataclass
class Subnets(Bootstrappable):
    # Output
    subnets: List[str] = field(init=False, default_factory=lambda: [])

    def bootstrap(self):
        """Find supported subnets.
        """
        super().bootstrap()
        ec2 = boto3.client("ec2")
        # Find the default VPC
        vpc_response = ec2.describe_vpcs(Filters=[{"Name": "isDefault", "Values": ["true"]}])
        default_vpc_id = vpc_response['Vpcs'][0]['VpcId']

        # Find all the subnets from default VPC
        subnet_response = ec2.describe_subnets(Filters=[{"Name": "vpc-id", "Values": [default_vpc_id]}])

        ec2_subnets = []
        for subnet in subnet_response['Subnets']:
            ec2_subnets.append(subnet['SubnetId'])

        # Try to create the subnet using all the available subnets
        mdb = boto3.client("memorydb")
        subnet_name = random_suffix_name("sub", 10)
        try:
            mdb.create_subnet_group(SubnetGroupName=subnet_name, Description='Determine valid subnets',
                                    SubnetIds=ec2_subnets)
        except mdb.exceptions.SubnetNotAllowedFault as ex:
            message = str(ex)
            exp_message = "Supported availability zones are "
            index = message.index(exp_message)
            # Format of the message is like "Supported availability zones are [us-east-1c, us-east-1d, us-east-1b]."
            valid_az = message[index + len(exp_message) + 1: len(message) - 2].split(", ")

            for subnet in subnet_response['Subnets']:
                if subnet['AvailabilityZone'] in valid_az:
                    self.subnets.append(subnet['SubnetId'])
            return

        # We were able to create subnet group using all the subnets, so delete the subnet group that was created
        self.subnets = ec2_subnets
        mdb.delete_subnet_group(SubnetGroupName=subnet_name)

    def cleanup(self):
        """Delete the secret.
        """
        super().cleanup()