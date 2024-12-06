package org.akrck02.valhalla.core.sdk.model.device

import kotlinx.serialization.Contextual
import kotlinx.serialization.SerialName
import org.akrck02.valhalla.core.sdk.model.user.User

/**
 * This class represents a user device logged in the platform
 * and the corresponding token
 */
@Suppress("unused")
data class Device(

    @SerialName("_id") @Contextual var id: String? = null,

    var user: User? = null,
    var address: String? = null,
    var userAgent: String? = null,
    var token: String? = null,
)