// handler/course_handler.go

package handler

import (
	"net/http"
	"strconv"

	"github.com/9500073161/coursemanagement/mysqldbmodels"
	"github.com/gin-gonic/gin"
)

// CourseHandler handles HTTP requests related to courses
type CourseHandler struct {
	Repo mysqldbmodels.CourseRepository
}

// NewCourseHandler initializes a new instance of CourseHandler
func NewCourseHandler(repo mysqldbmodels.CourseRepository) *CourseHandler {
	return &CourseHandler{Repo: repo}
}

// CreateCourse handles HTTP POST requests to create a new course
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var course mysqldbmodels.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.Createcourse(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(http.StatusCreated, course)
}

// GetCourseByID handles HTTP GET requests to retrieve a course by ID
func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	course, err := h.Repo.GetcourseByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

// UpdateCourse handles HTTP PUT requests to update a course
func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	var course mysqldbmodels.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course.ID = uint(id)
	if err := h.Repo.Updatecourse(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
		return
	}

	c.JSON(http.StatusOK, course)
}

// DeleteCourse handles HTTP DELETE requests to delete a course
func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	if err := h.Repo.Deletecourse(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
