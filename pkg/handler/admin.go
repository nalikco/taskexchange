package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) adminStatistics(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 3 {
		newErrorResponse(c, http.StatusForbidden, "access denied")
		return
	}

	users, err := h.services.Users.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	usersCount := 0
	performersCount := 0
	customersCount := 0
	usersBalance := 0

	for _, user := range users {
		usersCount++
		if user.Type == 1 {
			performersCount++
		}
		if user.Type == 2 {
			customersCount++
		}

		usersBalance += int(user.Balance)
	}

	activeTasksCount, err := h.services.Tasks.CountActive()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	activeOrdersCount, err := h.services.Orders.CountAllActive()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	completedOrders, err := h.services.Orders.GetAllCompleted()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	completedOrdersCount := 0
	completedOrdersPrice := 0
	completedOrdersTodayPrice := 0
	completedOrdersCurrentMonthPrice := 0
	currentYear, currentMonth, currentDay := time.Now().Date()

	for _, order := range completedOrders {
		completedOrdersCount++

		orderPrice := int(order.Task.CalculatePriceForOne())
		completedOrdersPrice += orderPrice

		orderYear, orderMonth, orderDay := order.CreatedAt.Date()

		if currentYear == orderYear && currentMonth == orderMonth {
			completedOrdersCurrentMonthPrice += orderPrice

			if currentDay == orderDay {
				completedOrdersTodayPrice += orderPrice
			}
		}

	}

	var fields = map[string]int{
		"users_count":                          usersCount,
		"performers_count":                     performersCount,
		"customers_count":                      customersCount,
		"users_balance":                        usersBalance,
		"active_tasks_count":                   activeTasksCount,
		"active_orders_count":                  activeOrdersCount,
		"completed_orders_count":               completedOrdersCount,
		"completed_orders_price":               completedOrdersPrice,
		"completed_orders_today_price":         completedOrdersTodayPrice,
		"completed_orders_current_month_price": completedOrdersCurrentMonthPrice,
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": fields,
	})
}
