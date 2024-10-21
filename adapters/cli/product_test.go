package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pedrojpx/hexagonal-example/adapters/cli"
	mock_application "github.com/pedrojpx/hexagonal-example/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "productTest"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expected := fmt.Sprintf("Product ID %s with name %s has been created with price %f and status %s", productMock.GetId(), productMock.GetName(), productMock.GetPrice(), productMock.GetStatus())
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expected, result)

	expected = fmt.Sprintf("Product %s has been enabled", productMock.GetName())
	result, err = cli.Run(service, "enable", productId, "", 0.0)
	require.Nil(t, err)
	require.Equal(t, expected, result)

	expected = fmt.Sprintf("Product %s has been disabled", productMock.GetName())
	result, err = cli.Run(service, "disable", productId, "", 0.0)
	require.Nil(t, err)
	require.Equal(t, expected, result)

	expected = fmt.Sprintf("Product %s has been disabled", productMock.GetName())
	result, err = cli.Run(service, "disable", productId, "", 0.0)
	require.Nil(t, err)
	require.Equal(t, expected, result)

	expected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s", productMock.GetId(), productMock.GetName(), productMock.GetPrice(), productMock.GetStatus())
	result, err = cli.Run(service, "", productId, "", 0.0)
	require.Nil(t, err)
	require.Equal(t, expected, result)
}
