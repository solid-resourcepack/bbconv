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

    private lateinit var entity: RenderedEntity

    override fun onCommand(sender: CommandSender, command: Command, label: String, args: Array<out String>): Boolean {
        val path = File(this.dataFolder, "baseconfig.json").toPath()
        val config = OpenModelLoader(path).load() ?: return false

        if (!::entity.isInitialized) {
            entity = RenderedEntity((sender as Player).location, config)
            entity.spawn()
            if (args.isNotEmpty()) {
                entity.getAnimationController().play(args[0])
            }
            return true
        }
        entity.getAnimationController().animate(0.10f)
        return true
    }

}