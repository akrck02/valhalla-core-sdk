package org.valhalla.core.sdk.error

/**
 * This enumeration represents the errors
 * captured and recognized by the service.
 */
@Suppress("unused")
enum class ErrorCode(val code: Int) {

    // 0 --> 999 | GENERAL ERRORS
    UnexpectedError(0),
    AccessDenied(1),
    NotImplemented(2),
    InvalidRequest(3),
    DatabaseError(4),
    InvalidId(5),
    NotFound(6),
    NotEnoughPermissions(6),
    InvalidPassword(7),
    InvalidToken(8),
    InvalidEmail(9),
    CannotCreateValidationCode(10),
    InvalidValidationCode(11),
    NothingChanged(12),

    // 1000 --> 1999 | USER ERRORS
    UserAlreadyExists(1000),
    UserAlreadyValidated(1001),
    UserNotAuthorized(1002),

    // 2000 --> 2999 | PROJECT ERRORS
    ProjectAlreadyExists(2000),

    // 3000 --> 3999 | TEAM ERRORS
    TeamAlreadyExists(3000),
    UserIsAlreadyMember(3001),
    UserIsNotMember(3002),

    // 4000 --> 4999 | DEVICE ERRORS
    DeviceNotFound(4000),
    CannotGenerateAuthToken(4001),
}