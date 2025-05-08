package io.github.solid.resourcepack.bbconv.api.animation.entity

import io.github.solid.resourcepack.bbconv.config.Animation

data class EntityAnimationData(
    val name: String,
    /**
     * Negative if the duration is indefinite/procedural
     */
    val duration: Float = -1f,
    val looped: Boolean = false,
    val loopDelay: Float = -1f,
) {
    companion object {
        @JvmStatic
        fun of(animation: Animation): EntityAnimationData {
            return EntityAnimationData(
                name = animation.name,
                duration = animation.length.toFloat(),
                looped = animation.loop,
                loopDelay = animation.loopDelay
            )
        }
    }
}
