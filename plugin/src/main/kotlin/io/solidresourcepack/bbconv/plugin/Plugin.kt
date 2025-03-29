package io.solidresourcepack.bbconv.plugin

import com.sun.tools.javac.tree.TreeInfo.args
import org.bukkit.plugin.java.JavaPlugin
import org.spongepowered.configurate.CommentedConfigurationNode
import org.spongepowered.configurate.gson.GsonConfigurationLoader
import java.io.File
import java.nio.file.Path
import java.nio.file.Paths


class Plugin : JavaPlugin() {

    override fun onEnable() {
        val path = File(this.dataFolder, "baseconfig.json").toPath()
        val loader = GsonConfigurationLoader.builder()
            .path(path) // or url(), or source/sink
            .build()

        val node = loader.load() // Load from file
        val config = node.get(BaseConfig::class.java) // Populate object

        println(config)
    }
}