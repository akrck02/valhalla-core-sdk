package org.akrck02.valhalla.core.sdk.model.device

import org.akrck02.valhalla.core.sdk.model.user.User

/**
 * This class represents a user device logged in the platform
 * and the corresponding token
 */
@Suppress("unused")
data class Device(
    val id: String?,
    var user: User? = null,
    var address: String? = null,
    var userAgent: String? = null,
    var token: String? = null
)