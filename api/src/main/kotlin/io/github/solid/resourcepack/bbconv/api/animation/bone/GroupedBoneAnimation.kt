package io.github.solid.resourcepack.bbconv.api.animation.bone

import org.bukkit.util.Transformation
import org.joml.Quaternionf
import org.joml.Vector3f

class GroupedBoneAnimation(
    private val animations: List<BoneAnimation>
) : BoneAnimation {
    override fun animate(context: BoneAnimationContext): Transformation {
        var result = Transformation(Vector3f(), Quaternionf(), Vector3f(), Quaternionf())
        animations.forEach {
            val transformation = it.animate(context)
            result = Transformation(
                result.translation.add(transformation.translation),
                result.leftRotation.add(transformation.leftRotation),
                result.scale.add(transformation.scale),
                Quaternionf()
            )
        }
        return result
    }
}