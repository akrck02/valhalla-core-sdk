import java.io.FileInputStream
import java.util.*

// Variables section
val localProperties: Properties = Properties().apply {
    load(FileInputStream(File(rootProject.rootDir, "local.properties")))
}

group = providers.gradleProperty("organization").getOrElse("")
version = providers.gradleProperty("valhalla.core.sdk.version").getOrElse("")

val jdkVersion: Int = providers.gradleProperty("jdk.version").getOrElse("").toInt()
val mavenServerName: String = localProperties.getProperty("maven.server.name")
val mavenServerUrl: String = localProperties.getProperty("maven.server.url")
val mavenServerUser: String = localProperties.getProperty("maven.server.user")
val mavenServerPassword: String = localProperties.getProperty("maven.server.password")

// Plugins section
plugins {
    id("org.jetbrains.kotlin.jvm") version "2.0.21"
    id("maven-publish")
    id("org.jetbrains.kotlin.plugin.serialization") version "2.1.0"
}

// Repositories
val selfHostedRepository = repositories.maven {
    name = mavenServerName
    url = uri(mavenServerUrl)
    credentials {
        username = mavenServerUser
        password = mavenServerPassword
    }
    isAllowInsecureProtocol = true
}

repositories {
    mavenCentral()
    selfHostedRepository
//    mavenLocal()
}

// Testing section
tasks.test {
    useJUnit()
}

// Compilation section
kotlin {
    jvmToolchain(jdkVersion)
}

java {
    withSourcesJar()
    withJavadocJar()
}

// Publishing section
publishing {
    repositories {
        mavenLocal()
        selfHostedRepository
    }
    publications.withType<MavenPublication> {
        from(components["java"])
    }
}

// Dependency section
dependencies {

    // Tests
    testImplementation("org.jetbrains.kotlin:kotlin-test")

    // Kotlin coroutines
    implementation("org.mongodb:mongodb-driver-kotlin-coroutine:5.2.1")

    // Serialization
    implementation("org.jetbrains.kotlinx:kotlinx-serialization-core:1.5.1")
    implementation("org.jetbrains.kotlinx:kotlinx-serialization-json:1.7.3")
    implementation("org.mongodb:bson:5.2.1")
}




