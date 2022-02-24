res, err := rm.validateACLNeedsUpdate(latest)

if err != nil || res!= nil{
    return res, err
}