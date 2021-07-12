package ginrestaurant

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/component/appctx"
	"fooddelivery/module/restaurant/business"
	"fooddelivery/module/restaurant/model"
	restaurantstorage "fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

type fakeListStore struct{}

func (fakeListStore) ListDataWithCondition(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	return []restaurantmodel.Restaurant{
		{
			SQLModel: common.SQLModel{ID: 1},
			Name:     "AA",
			Address:  "BB",
		},
	}, nil
}

func ListRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Process()

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appContext.GetMainDBConnection())
		business := restaurantbusiness.NewlistRestaurantBusiness(store)

		result, err := business.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(common.ErrInternal(err))
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
