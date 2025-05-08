package io.github.solid.resourcepack.bbconv.config

import org.spongepowered.configurate.gson.GsonConfigurationLoader
import java.nio.file.Path

class OpenModelLoader(private val path: Path) {
    fun load(): OpenModelConfig? {
        val loader = GsonConfigurationLoader.builder()
            .path(path)
            .build()

        val node = loader.load()
        return node.get(OpenModelConfig::class.java)
    }
}