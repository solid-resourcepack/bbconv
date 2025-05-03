package io.github.solid.resourcepack.bbconv.api.animation.entity

import io.github.solid.resourcepack.bbconv.api.animation.bone.BoneAnimation
import io.github.solid.resourcepack.bbconv.api.animation.bone.KeyFramedBoneAnimation
import io.github.solid.resourcepack.bbconv.config.Animation

typealias EntityAnimation = MutableMap<String, BoneAnimation>

object EntityAnimations {
    fun of(animation: Animation): EntityAnimation {
        return animation.animators.associate { animator -> animator.bone to KeyFramedBoneAnimation(animator) }
            .toMutableMap()
    }
}

