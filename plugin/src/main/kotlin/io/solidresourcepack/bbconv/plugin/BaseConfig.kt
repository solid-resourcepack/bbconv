package io.solidresourcepack.bbconv.plugin

import org.spongepowered.configurate.objectmapping.ConfigSerializable
import org.spongepowered.configurate.objectmapping.meta.Setting

@ConfigSerializable
data class Bone(
    val id: String = "",

    val model: String = "",

    val children: List<Bone> = emptyList(),

    val origin: List<Float> = emptyList(),

    val visible: Boolean = false,

    val scale: Float = 0f,
)

@ConfigSerializable
data class Animation(
    val loop: Boolean = false,

    val length: Double = 0.0,

    @Setting("start_delay")
    val startDelay: Float = 0f,

    @Setting("loop_delay")
    val loopDelay: Float = 0f,

    val name: String = "",

    val animators: List<Animator> = emptyList(),
)

@ConfigSerializable
data class Animator(
    val name: String = "",

    val bone: String = "",

    val position: List<PositionKeyframe> = emptyList(),

    val rotation: List<RotationKeyframe> = emptyList(),

    val scale: List<ScaleKeyframe> = emptyList(),
)

@ConfigSerializable
data class PositionKeyframe(
    val time: Float = 0f,

    val position: Vector = Vector(0f, 0f, 0f),

    val interpolation: String = "",
)

@ConfigSerializable
data class RotationKeyframe(
    val time: Float = 0f,

    @Setting("left_rotation")
    val leftRotation: Quaternion = Quaternion(0f, 0f, 0f, 0f),

    @Setting("right_rotation")
    val rightRotation: Quaternion = Quaternion(0f, 0f, 0f, 0f),

    val interpolation: String = "",
)

@ConfigSerializable
data class ScaleKeyframe(
    val time: Float,

    val scale: Vector = Vector(0f, 0f, 0f),

    val interpolation: String = "",
)

@ConfigSerializable
data class Point(
    val x: Float = 0f,

    val y: Float = 0f,

    val z: Float = 0f,
)

@ConfigSerializable
data class Vector(
    val x: Float = 0f,

    val y: Float = 0f,

    val z: Float = 0f,
)

@ConfigSerializable
data class Quaternion(
    val w: Float = 0f,

    val x: Float = 0f,

    val y: Float = 0f,

    val z: Float = 0f,
)

@ConfigSerializable
data class BaseConfig(
    val name: String = "",

    @Setting("bone_tree")
    val boneTree: List<Bone> = emptyList(),

    val animations: List<Animation> = emptyList()
)
