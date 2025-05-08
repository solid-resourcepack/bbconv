package io.github.solid.resourcepack.bbconv.api.animation.keyframe

import io.github.solid.resourcepack.bbconv.util.Interpolation

data class Keyframe<T>(
    val data: T,
    val interpolation: Interpolation,
)