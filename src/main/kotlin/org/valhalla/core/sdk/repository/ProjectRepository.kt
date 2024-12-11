package org.valhalla.core.sdk.repository

import org.valhalla.core.sdk.model.exception.ServiceException
import org.valhalla.core.sdk.model.project.Project

/** This interface represents the data access layer for the projects.*/
@Suppress("unused")
interface ProjectRepository {

    /**
     * Register a new [Project] returning the id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun register(project: Project?): String

    /**
     * Get a [Project] by id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun get(id: String?): Project

    /**
     * Get a [List] of the [Project]s a user belongs to by user id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun getAllByMember(userId: String?): List<Project>

    /**
     * Update the [Project] data by id and new project data.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun update(id: String?, project: Project?)

    /**
     * Delete a [Project] by id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun delete(id: String?)

}