package io.solidresourcepack.bbconv.plugin

import io.github.solid.resourcepack.bbconv.api.entity.RenderedEntity
import io.github.solid.resourcepack.bbconv.config.OpenModelLoader
import org.bukkit.Bukkit
import org.bukkit.command.Command
import org.bukkit.command.CommandSender
import org.bukkit.entity.Player
import org.bukkit.plugin.java.JavaPlugin
import org.bukkit.scheduler.BukkitTask
import java.io.File


class Plugin : JavaPlugin() {

    override fun onEnable() {
        this.getCommand("t")?.setExecutor(this)
    }

    private var entity: RenderedEntity? = null

    override fun onCommand(sender: CommandSender, command: Command, label: String, args: Array<out String>): Boolean {
        val path = File(this.dataFolder, "baseconfig.json").toPath()
        val config = OpenModelLoader(path).load() ?: return false

        if (entity == null) {
            entity = RenderedEntity((sender as Player).location, config)
            entity!!.spawn()
            if (args.isNotEmpty()) {
                val type = args[0]
                val animation = entity!!.getAnimationController().get(type) ?: return false
                entity!!.getAnimationController().play(type)
                var elapsed = 0f
                var task: BukkitTask? = null
                task = Bukkit.getScheduler().runTaskTimer(this, Runnable {
                    entity!!.getAnimationController().animate(0.05f)
                    elapsed += 0.05f
                    if (elapsed > animation.duration) {
                        task?.cancel()
                        entity = null
                    }
                }, 5L, 1L)
            }
            return true
        }
        return true
    }

}