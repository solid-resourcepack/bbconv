plugins {
    kotlin("jvm") version "2.1.10"
    id("com.gradleup.shadow") version "8.3.3"
}

group = "io.solid-resourcepack.bbconv"
version = "0.1.0"

allprojects {
    repositories {
        mavenCentral()
        maven {
            name = "papermc"
            url = uri("https://repo.papermc.io/repository/maven-public/")
        }
    }
}

subprojects {

    apply {
        plugin("org.jetbrains.kotlin.jvm")
        plugin("com.gradleup.shadow")
    }

    dependencies {
        testImplementation(kotlin("test"))
        implementation("org.spongepowered:configurate-extra-kotlin:4.1.2")
        implementation("org.spongepowered:configurate-gson:4.1.2")
        compileOnly("io.papermc.paper:paper-api:1.21.4-R0.1-SNAPSHOT")
    }

    tasks.test {
        useJUnitPlatform()
    }

    kotlin {
        jvmToolchain(21)
    }
}

