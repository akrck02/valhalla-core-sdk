package org.akrck02.valhalla.core.sdk.model.device

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable
import org.akrck02.valhalla.core.sdk.model.user.User

/**
 * This class represents a user device logged in the platform
 * and the corresponding token
 */
@Suppress("unused")
@Serializable
data class Device(
    @SerialName("_id") var id: String? = null,
    var user: User? = null,
    var address: String? = null,
    var userAgent: String? = null,
    var token: String? = null
)