package app

//
//func TestApp_RegisterRoute(t *testing.T) {
//
//}
//
//func TestApp_GetRoute(t *testing.T) {
//	// Setup
//	e := echo.New()
//	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	app := App{
//		UserRepo: nil,
//		URLRepo:  nil,
//		JWTGen:   nil,
//	}
//
//	// Assertions
//	if assert.NoError(t, h.createUser(c)) {
//		assert.Equal(t, http.StatusCreated, rec.Code)
//		assert.Equal(t, userJSON, rec.Body.String())
//	}
//}
