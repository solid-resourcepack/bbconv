package io.github.solid.resourcepack.bbconv.api.entity

import io.github.solid.resourcepack.bbconv.config.OpenModelConfig
import org.bukkit.Location
import java.util.*

data class RenderedEntity(
    val location: Location,
    val config: OpenModelConfig,
    val id: UUID = UUID.randomUUID(),
) {
    private val renderedBones = mutableMapOf<String, RenderedBone>()

    init {
        location.pitch = 0f
        config.getBoneMap().filter { it.value.visible }.forEach { (key, bone) ->
            renderedBones[key] = RenderedBone(this, bone)
        }
    }

    fun spawn() {
        renderedBones.values.forEach { it.spawn() }
    }
}