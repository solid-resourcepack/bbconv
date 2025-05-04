package io.github.solid.resourcepack.bbconv.api.entity

import io.github.solid.resourcepack.bbconv.config.Bone
import org.bukkit.Material
import org.bukkit.NamespacedKey
import org.bukkit.entity.EntityType
import org.bukkit.entity.ItemDisplay
import org.bukkit.inventory.ItemStack
import org.bukkit.persistence.PersistentDataType
import org.bukkit.util.Transformation
import org.joml.Quaternionf
import org.joml.Vector3f
import java.util.*

class RenderedBone(
    private val entity: RenderedEntity,
    val bone: Bone,
    val children: MutableList<RenderedBone> = mutableListOf()
) {

    private lateinit var display: ItemDisplay
    private val initialTransformation = Transformation(
        Vector3f(bone.origin[0], bone.origin[1], bone.origin[2]),
        bone.leftRotation.toQuaternionf(),
        Vector3f(bone.scale),
        Quaternionf(),
    )

    fun getInitialTransformation(): Transformation {
        return Transformation(
            Vector3f(initialTransformation.translation), Quaternionf(initialTransformation.leftRotation),
            Vector3f(initialTransformation.scale), Quaternionf()
        )
    }

    fun spawn() {
        if (::display.isInitialized) {
            return
        }
        display = entity.location.world.spawnEntity(entity.location, EntityType.ITEM_DISPLAY) as ItemDisplay
        display.isInvisible = !bone.visible
        display.itemDisplayTransform = ItemDisplay.ItemDisplayTransform.FIXED
        val stack = ItemStack(Material.BONE)
        stack.editMeta {
            it.itemModel = NamespacedKey.fromString(bone.model.replaceFirst("item/", ""))
        }
        display.setItemStack(stack)
        display.transformation = getInitialTransformation()
        display.interpolationDuration = 1
        markDisplay()
        return
    }

    // Mark our current display to be a certain bone of a certain entity
    private fun markDisplay() {
        display.persistentDataContainer.set(
            NamespacedKey("bbconv", "bone_id"), PersistentDataType.STRING, bone.id
        )
        display.persistentDataContainer.set(
            NamespacedKey(
                "bbconv", "entity_id"
            ), PersistentDataType.STRING, entity.id.toString()
        )
    }

    fun getDisplay(): ItemDisplay {
        return display
    }

    companion object {

        /**
         * Find the rendered bone of an entity in the world
         */
        @JvmStatic
        fun find(entity: RenderedEntity, bone: Bone): Optional<RenderedBone> {
            entity.location.world.getNearbyEntitiesByType(ItemDisplay::class.java, entity.location, 20.0)
                .forEach { display ->
                    val boneId = display.persistentDataContainer.get(
                        NamespacedKey("bbconv", "bone_id"),
                        PersistentDataType.STRING
                    )
                    val entityId = display.persistentDataContainer.get(
                        NamespacedKey("bbconv", "entity_id"),
                        PersistentDataType.STRING
                    )
                    if (entityId == entity.id.toString() && boneId == bone.id) {
                        val rendered = RenderedBone(entity, bone)
                        rendered.display = display
                        return Optional.of(rendered)
                    }
                }
            return Optional.empty()
        }
    }
}