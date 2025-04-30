package io.solidresourcepack.bbconv.plugin

import io.github.solid.resourcepack.bbconv.api.entity.RenderedEntity
import io.github.solid.resourcepack.bbconv.config.OpenModelLoader
import org.bukkit.command.Command
import org.bukkit.command.CommandSender
import org.bukkit.entity.Player
import org.bukkit.plugin.java.JavaPlugin
import java.io.File


class Plugin : JavaPlugin() {

    override fun onEnable() {
        this.getCommand("t")?.setExecutor(this)
    }

    override fun onCommand(sender: CommandSender, command: Command, label: String, args: Array<out String>): Boolean {
        val path = File(this.dataFolder, "baseconfig.json").toPath()
        val config = OpenModelLoader(path).load() ?: return false

        val entity = RenderedEntity((sender as Player).location, config)
        entity.spawn()

        return true
    }

}