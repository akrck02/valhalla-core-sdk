package org.valhalla.core.sdk.repository

import org.valhalla.core.sdk.model.device.Device
import org.valhalla.core.sdk.model.exception.ServiceException

/** This interface represents the data access layer for the user devices. */
@Suppress("unused")
interface DeviceRepository {

    /**
     * Register a new [Device] for a user given a user id, returning the auth token.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun register(userId: String?, device: Device?): String

    /**
     * Get a [Device] for a user given the id and a user id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun get(userId: String?, id: String?): Device

    /**
     * Get a [Device] by user and current auth token.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun getByAuth(userId: String?, token: String?): Device

    /**
     * Get the [List] of [Device]s belonging to a user.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun getAll(userId: String?): List<Device>

}