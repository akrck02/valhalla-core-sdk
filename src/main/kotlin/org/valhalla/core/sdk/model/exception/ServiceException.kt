package org.valhalla.core.sdk.model.exception

import kotlinx.serialization.Serializable
import org.valhalla.core.sdk.error.ErrorCode
import org.valhalla.core.sdk.model.http.HttpStatusCode

/** This class represents an exception in the service. */
@Suppress("unused")
@Serializable
class ServiceException(
    val status: HttpStatusCode = HttpStatusCode.InternalServerError,
    val code: ErrorCode? = ErrorCode.UnexpectedError,
    override val message: String = "Unexpected error"
) : Exception(message) {

    override fun toString(): String {
        return "status: $status, code: $code, message: $message"
    }

}
