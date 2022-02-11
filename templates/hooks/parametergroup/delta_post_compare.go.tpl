if delta.DifferentAt("Spec.Description") {
    common.RemoveFromDelta(delta, "Spec.Description")
}

if delta.DifferentAt("Spec.Name") {
    common.RemoveFromDelta(delta, "Spec.Name")
}

if delta.DifferentAt("Spec.Family") {
    common.RemoveFromDelta(delta, "Spec.Family")
}