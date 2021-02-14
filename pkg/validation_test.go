package pkg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsPhoneValid(t *testing.T) {
	require.True(t, IsPhoneValid("+79045195421"))
	require.True(t, IsPhoneValid("+992928612100"))
	require.True(t, IsPhoneValid("+992930004800"))
	require.True(t, IsPhoneValid("+992-92-704-21-04"))
	require.True(t, IsPhoneValid("+992-93-(852)-57-75"))

	require.False(t, IsPhoneValid("a31"))
	require.False(t, IsPhoneValid("43231bd174"))
}

func TestIsEmailValid(t *testing.T) {
	require.True(t, IsEmailValid("hakimov@hotmail.com"))
	require.True(t, IsEmailValid("jakha.m.jm.m2000@mail.ru"))
	require.True(t, IsEmailValid("mr.adams@gmail.com"))

	require.False(t, IsEmailValid("h@mail."))
	require.False(t, IsEmailValid("superuser@supermail.c"))
}
