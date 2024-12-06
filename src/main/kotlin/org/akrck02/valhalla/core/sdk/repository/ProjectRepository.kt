package org.akrck02.valhalla.core.sdk.repository

import org.akrck02.valhalla.core.sdk.model.project.Project

/**
 * This interface represents the data access layer
 * for the team collection.
 */
@Suppress("unused")
interface ProjectRepository {

    /**
     * Register a new project
     * @param project The project to register
     */
    suspend fun register(project: Project?)

    /**
     * Get a project by id
     * @param id The project id
     * @return The requested project
     */
    suspend fun get(id: String?): Project?

    /**
     * Get all the projects a user belongs to
     * @param userId The user id
     * @return the list of projects a user belongs to
     */
    suspend fun getAllByMember(userId: String?): List<Project>

    /**
     * Update the project data
     * @param id The project id
     * @param project The new project data
     */
    suspend fun update(id: String?, project: Project?)

    /**
     * Delete a project
     * @param id The project id
     */
    suspend fun delete(id: String?)

}