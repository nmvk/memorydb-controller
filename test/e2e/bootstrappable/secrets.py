

from dataclasses import dataclass, field

from acktest.bootstrapping import Bootstrappable
from acktest import resources
from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s


@dataclass
class Secret(Bootstrappable):
    # Inputs
    name: str

    def bootstrap(self):
        """Create a new secret.
        """
        super().bootstrap()
        k8s.create_opaque_secret("default", self.name, "password", random_suffix_name("password", 32))

    def cleanup(self):
        """Delete the secret.
        """
        super().cleanup()
        k8s.delete_secret("default", self.name)