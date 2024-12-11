package org.valhalla.core.sdk.model.device

import kotlinx.serialization.Serializable

/**
 * This class represents a user device logged in the platform
 * and the corresponding token.
 */
@Suppress("unused")
@Serializable
data class Device(
    var address: String? = null,
    var userAgent: String? = null,
    var token: String? = null
)