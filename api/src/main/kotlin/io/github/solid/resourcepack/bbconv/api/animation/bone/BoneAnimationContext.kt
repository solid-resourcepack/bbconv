package io.github.solid.resourcepack.bbconv.api.animation.bone

import io.github.solid.resourcepack.bbconv.api.entity.RenderedBone
import io.github.solid.resourcepack.bbconv.api.entity.RenderedEntity

/**
 * Context of an active BoneAnimation.
 * This holds the bone itself, the current entity and the
 * elapsed time for the current animation and the initial state (transformation property pre mutation) of the bone.
 *
 * Mutating the transformation property will apply this mutation to the [RenderedBone]
 */
data class BoneAnimationContext(
    val entity: RenderedEntity,
    val elapsedTime: Float,
    val self: RenderedBone,
)
