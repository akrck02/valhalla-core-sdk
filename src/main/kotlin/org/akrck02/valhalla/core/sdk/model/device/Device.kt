package org.akrck02.valhalla.core.sdk.model.device

import org.akrck02.valhalla.core.sdk.model.user.User

/**
 * This class represents a user device logged in the platform
 * and the corresponding token
 */
@Suppress("unused")
data class Device(
    var id: String?,
    var user: User?,
    var address: String?,
    var userAgent: String?,
    var token: String?
)