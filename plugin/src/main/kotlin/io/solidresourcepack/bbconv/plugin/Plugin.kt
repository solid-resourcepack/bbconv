package io.solidresourcepack.bbconv.plugin

import com.sun.tools.javac.tree.TreeInfo.args
import org.bukkit.Location
import org.bukkit.Material
import org.bukkit.NamespacedKey
import org.bukkit.command.Command
import org.bukkit.command.CommandSender
import org.bukkit.entity.EntityType
import org.bukkit.entity.ItemDisplay
import org.bukkit.entity.Player
import org.bukkit.inventory.ItemStack
import org.bukkit.plugin.java.JavaPlugin
import org.bukkit.util.Transformation
import org.joml.Quaternionf
import org.joml.Vector3f
import org.spongepowered.configurate.CommentedConfigurationNode
import org.spongepowered.configurate.gson.GsonConfigurationLoader
import java.io.File
import java.nio.file.Path
import java.nio.file.Paths


class Plugin : JavaPlugin() {

    override fun onEnable() {
        this.getCommand("t")?.setExecutor(this)
    }

    override fun onCommand(sender: CommandSender, command: Command, label: String, args: Array<out String>): Boolean {
        val path = File(this.dataFolder, "baseconfig.json").toPath()
        val loader = GsonConfigurationLoader.builder()
            .path(path)
            .build()

        val node = loader.load()
        val config = node.get(BaseConfig::class.java) ?: return false

        renderTree(config.boneTree, (sender as Player).location)

        return true
    }

    fun renderTree(bones: List<Bone>, loc: Location) {
        for (bone in bones) {
            println(bone.id)
            if (bone.scale != 0f) {
                boneToItemDisplay(bone, loc)
            }
            renderTree(bone.children, loc)
        }
    }

    fun boneToItemDisplay(bone: Bone, loc: Location) {
        val loc1 = loc.toBlockLocation()
        loc1.yaw = 0f
        loc1.pitch = 0f
        val entity = loc.world.spawnEntity(loc1, EntityType.ITEM_DISPLAY) as ItemDisplay

        val stack = ItemStack(Material.BONE)
        stack.editMeta {
            it.itemModel = NamespacedKey.fromString(bone.model.replaceFirst("item/", ""))
        }

        entity.isInvisible = false
        entity.setItemStack(stack)
        entity.transformation = Transformation(
            Vector3f(bone.origin[0], bone.origin[1], bone.origin[2]).div(16f),
            Quaternionf(),
            Vector3f(1 / bone.scale),
            Quaternionf(),
        )
    }

}