package middlewares

//func ACL(permissionToCheck string) echo.MiddlewareFunc {
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			user, ok := c.Get("user").(*serializers.LoggedInUser)
//			if !ok {
//				return c.JSON(http.StatusInternalServerError, msgutil.NewRestResp("no logged-in user found", nil))
//			}
//
//			if user.HasPermission(permissionToCheck) {
//				return next(c)
//			}
//
//			return c.JSON(http.StatusForbidden, msgutil.NewRestResp("access forbidden", nil))
//		}
//	}
//}
