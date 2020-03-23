package wallet

import "testing"

func TestWallet(t *testing.T) {

	t.Run("should deposit coins", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(&wallet, Bitcoin(10), t)
	})

	t.Run("should withdraw coins", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(7))

		assertBalance(&wallet, Bitcoin(3), t)
		assertNoError(err, t)
	})

	t.Run("should error on insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(&wallet, startingBalance, t)
		assertError(err, InsufficientFundsErr, t)

	})
}

func assertBalance(wallet *Wallet, expected Bitcoin, t *testing.T) {
	t.Helper()
	actual := wallet.Balance()
	if actual != expected {
		t.Errorf("Expected balance: %s, actual balance: %s", expected, actual)
	}
}

func assertError(err error, expected error, t *testing.T) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected error, but none was returned")
	}

	if err != expected {
		t.Errorf("Expected error %s, actual error %s", expected, err.Error())
	}
}

func assertNoError(err error, t *testing.T) {
	t.Helper()

	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}
}
