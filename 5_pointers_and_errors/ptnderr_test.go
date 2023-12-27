package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("test deposit on wallet", func(t *testing.T) {
		wallet := Wallet{}
		bitcoin := Bitcoin(10)

		wallet.Deposit(bitcoin)

		assertBalance(t, wallet, Bitcoin(10.0))
	})

	t.Run("test withdraw from wallet", func(t *testing.T) {
		wallet := Wallet{40.0}
		bitcoin := Bitcoin(10)

		err := wallet.Withdraw(bitcoin)

		assertBalance(t, wallet, Bitcoin(30.0))
		assertNoError(t, err)
	})


	t.Run("test withdraw from wallet with inssuficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(40.0)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100.0))
		
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
	

}


func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("%#v got %s and want %s", wallet, got, want)
	}	
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted an error but got nothing")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatalf("expected nothing but got %s", got)
	}
}
