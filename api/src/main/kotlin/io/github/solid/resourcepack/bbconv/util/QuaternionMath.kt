package io.github.solid.resourcepack.bbconv.util

import org.joml.Quaternionf

object QuaternionMath {
    fun delta(from: Quaternionf, to: Quaternionf): Quaternionf {
        val invFrom = Quaternionf(from).invert()
        return Quaternionf(to).mul(invFrom)
    }
}