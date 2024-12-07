package org.akrck02.valhalla.core.sdk.model.device

import kotlinx.serialization.Serializable
import org.akrck02.valhalla.core.sdk.model.user.User

/**
 * This class represents a user device logged in the platform
 * and the corresponding token
 */
@Suppress("unused")
@Serializable
data class Device(
    var id: String? = null,
    var owner: User? = null,
    var address: String? = null,
    var userAgent: String? = null,
    var token: String? = null
)