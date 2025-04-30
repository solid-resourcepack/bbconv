package io.github.solid.resourcepack.bbconv.util

import org.joml.Quaternionf
import org.joml.Vector3f

enum class Interpolation {
    LINEAR {
        override fun interpolate(t: Float, vararg v: Vector3f): Vector3f {
            return InterpolationUtils.lerp(v[0], v[1], t)
        }

        override fun interpolate(t: Float, vararg q: Quaternionf): Quaternionf {
            return InterpolationUtils.slerp(q[0], q[1], t)
        }
    },
    CATMULLROM {
        override fun interpolate(t: Float, vararg v: Vector3f): Vector3f {
            return InterpolationUtils.catmullRom(v[0], v[1], v[2], v[3], t)
        }

        override fun interpolate(t: Float, vararg q: Quaternionf): Quaternionf {
            return InterpolationUtils.quatCatmullRom(q[0], q[1], q[2], q[3], t)
        }
    };

    abstract fun interpolate(t: Float, vararg v: Vector3f): Vector3f
    abstract fun interpolate(t: Float, vararg q: Quaternionf): Quaternionf

    companion object {
        @JvmStatic
        fun fromString(interpolation: String): Interpolation {
            return valueOf(interpolation.uppercase())
        }
    }
}


object InterpolationUtils {

    fun lerp(a: Vector3f, b: Vector3f, t: Float): Vector3f {
        return Vector3f().set(a).lerp(b, t)
    }

    fun slerp(a: Quaternionf, b: Quaternionf, t: Float): Quaternionf {
        return Quaternionf(a).slerp(b, t)
    }

    fun catmullRom(p0: Vector3f, p1: Vector3f, p2: Vector3f, p3: Vector3f, t: Float): Vector3f {
        val t2 = t * t
        val t3 = t2 * t
        return Vector3f(
            0.5f * ((2f * p1.x) + (-p0.x + p2.x) * t + (2f * p0.x - 5f * p1.x + 4f * p2.x - p3.x) * t2 + (-p0.x + 3f * p1.x - 3f * p2.x + p3.x) * t3),

            0.5f * ((2f * p1.y) + (-p0.y + p2.y) * t + (2f * p0.y - 5f * p1.y + 4f * p2.y - p3.y) * t2 + (-p0.y + 3f * p1.y - 3f * p2.y + p3.y) * t3),

            0.5f * ((2f * p1.z) + (-p0.z + p2.z) * t + (2f * p0.z - 5f * p1.z + 4f * p2.z - p3.z) * t2 + (-p0.z + 3f * p1.z - 3f * p2.z + p3.z) * t3)
        )
    }

    // fake quaternion catmullrom (optional smooth approximation)

    fun quatCatmullRom(q0: Quaternionf, q1: Quaternionf, q2: Quaternionf, q3: Quaternionf, t: Float): Quaternionf {
        val slerp1 = Quaternionf(q1).slerp(q2, t)
        val slerp2 = Quaternionf(q0).slerp(q3, t)
        return slerp1.slerp(slerp2, 2f * t * (1f - t)) // not perfect but smoother than pure slerp
    }
}