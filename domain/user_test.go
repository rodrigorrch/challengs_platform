package domain_test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/rodrigorrch/challengs_platform/domain"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	_, err := domain.NewUser("Rodrigo", "rodrigoxrodrigoxxxgmail.com", "12345678")
	require.Error(t, err)

	_, err = domain.NewUser("Rodrigo", "rodrigoxrodrigoxxxgmail.com", "")
	require.Error(t, err)

	_, err = domain.NewUser("", "rodrigoxrodrigoxxxgmail.com", "")
	require.Error(t, err)

	user, err := domain.NewUser("Rodrigo", "rodrigoxrodrigoxxx@gmail.com", "12345678")
	require.Nil(t, err)

	govalidator.IsUUIDv4(user.Token)

}
