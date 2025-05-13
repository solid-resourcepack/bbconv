package io.github.solid.resourcepack.bbconv.api.animation.bone

import org.bukkit.util.Transformation

/**
 * Produces a Transformation based on the current animation context
 */
fun interface BoneAnimation {
    fun animate(context: BoneAnimationContext): Transformation
}