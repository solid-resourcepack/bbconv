package io.github.solid.resourcepack.bbconv.util

import org.joml.Quaternionf
import org.joml.Vector3f

enum class Interpolation {
    LINEAR {
        override fun <T> interpolate(
            t: Float,
            functions: InterpolationFunctions<T>,
            vararg v: T
        ): T {
            val first = if (v.size > 2) v[1] else v[0]
            val second = if (v.size > 2) v[2] else v[1]
            return functions.linear(first, second, t)
        }
    },

    CATMULLROM {
        override fun <T> interpolate(
            t: Float,
            functions: InterpolationFunctions<T>,
            vararg v: T
        ): T {
            return functions.catmullRom(v[0], v[1], v[1], v[3], t)
        }
    };

    abstract fun <T> interpolate(
        t: Float,
        functions: InterpolationFunctions<T>,
        vararg v: T
    ): T

    companion object {
        @JvmStatic
        fun fromString(interpolation: String): Interpolation {
            return valueOf(interpolation.uppercase())
        }
    }
}

class Interpolator<T>(
    private val interpolation: Interpolation,
    clazz: Class<T>
) {
    private val functions: InterpolationFunctions<T> = InterpolationFunctions.of(clazz)

    fun interpolate(t: Float, vararg values: T): T {
        return interpolation.interpolate(t, functions, *values)
    }
}

interface InterpolationFunctions<T> {
    fun linear(a: T, b: T, t: Float): T
    fun catmullRom(p0: T, p1: T, p2: T, p3: T, t: Float): T

    companion object {
        fun <T> of(clazz: Class<T>): InterpolationFunctions<T> {
            return when (clazz) {
                Vector3f::class.java -> VectorInterpolationFunctions as InterpolationFunctions<T>
                Quaternionf::class.java -> QuaternionInterpolationFunctions as InterpolationFunctions<T>
                else -> throw IllegalArgumentException("Unsupported type: $clazz")
            }
        }
    }
}

object VectorInterpolationFunctions : InterpolationFunctions<Vector3f> {
    override fun linear(a: Vector3f, b: Vector3f, t: Float): Vector3f {
        return Vector3f().set(a).lerp(b, t)
    }

    override fun catmullRom(p0: Vector3f, p1: Vector3f, p2: Vector3f, p3: Vector3f, t: Float): Vector3f {
        val t2 = t * t
        val t3 = t2 * t
        return Vector3f(
            0.5f * ((2f * p1.x) + (-p0.x + p2.x) * t + (2f * p0.x - 5f * p1.x + 4f * p2.x - p3.x) * t2 + (-p0.x + 3f * p1.x - 3f * p2.x + p3.x) * t3),
            0.5f * ((2f * p1.y) + (-p0.y + p2.y) * t + (2f * p0.y - 5f * p1.y + 4f * p2.y - p3.y) * t2 + (-p0.y + 3f * p1.y - 3f * p2.y + p3.y) * t3),
            0.5f * ((2f * p1.z) + (-p0.z + p2.z) * t + (2f * p0.z - 5f * p1.z + 4f * p2.z - p3.z) * t2 + (-p0.z + 3f * p1.z - 3f * p2.z + p3.z) * t3)
        )
    }
}

object QuaternionInterpolationFunctions : InterpolationFunctions<Quaternionf> {
    override fun linear(a: Quaternionf, b: Quaternionf, t: Float): Quaternionf {
        return Quaternionf(a).slerp(b, t)
    }

    override fun catmullRom(
        p0: Quaternionf,
        p1: Quaternionf,
        p2: Quaternionf,
        p3: Quaternionf,
        t: Float
    ): Quaternionf {
        val slerp1 = Quaternionf(p1).slerp(p2, t)
        val slerp2 = Quaternionf(p0).slerp(p3, t)
        return slerp1.slerp(slerp2, 2f * t * (1f - t)) // not perfect but smoother than pure slerp
    }

}