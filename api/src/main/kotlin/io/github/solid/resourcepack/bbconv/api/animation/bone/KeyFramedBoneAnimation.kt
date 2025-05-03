package io.github.solid.resourcepack.bbconv.api.animation.bone

import io.github.solid.resourcepack.bbconv.api.animation.keyframe.Timeline
import io.github.solid.resourcepack.bbconv.config.Animator
import org.bukkit.util.Transformation
import org.joml.Quaternionf
import org.joml.Vector3f

/**
 * Represents a keyframe based bone animation (being read from configs or other non-procedural sources)
 */
class KeyFramedBoneAnimation(
    private val animator: Animator
) : BoneAnimation {
    override fun animate(context: BoneAnimationContext): Transformation {
        val scale = if (animator.scale.size >= 2) Timeline.ofScale(animator.scale)
            .interpolate(context.elapsedTime) else Vector3f()
        val position = if (animator.position.size >= 2) Timeline.ofPosition(animator.position)
            .interpolate(context.elapsedTime) else Vector3f()
        val rotation = if (animator.rotation.size >= 2) Timeline.ofRotation(animator.rotation)
            .interpolate(context.elapsedTime) else Quaternionf()
        return Transformation(position, rotation, Vector3f(), Quaternionf())
    }
}