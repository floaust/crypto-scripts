package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// gcdCmd represents the gcd command
var gcdCmd = &cobra.Command{
	Use:   "gcd",
	Short: "greatest common divisor",
	Long: `The greatest common divisor (GCD) of two nonzero integers n and a is the greatest positive integer d such that d is a divisor of both n and a; that is, there are integers e and f such that n = de and a = df, and d is the largest such integer. The GCD of n and a is generally denoted gcd(n, a).[9]

This definition also applies when one of n and a is zero. In this case, the GCD is the absolute value of the non zero integer: gcd(n, 0) = gcd(0, n) = |n|. This case is important as the terminating step of the Euclidean algorithm.

The above definition cannot be used for defining gcd(0, 0), since 0 × n = 0, and zero thus has no greatest divisor. However, zero is its own greatest divisor if greatest is understood in the context of the divisibility relation, so gcd(0, 0) is commonly defined as 0. This preserves the usual identities for GCD, and in particular Bézout's identity, namely that gcd(n, a) generates the same ideal as {n, a}.[10][11][12] This convention is followed by many computer algebra systems.[13] Nonetheless, some authors leave gcd(0, 0) undefined.[14]

The GCD of n and b is their greatest positive common divisor in the preorder relation of divisibility. This means that the common divisors of n and a are exactly the divisors of their GCD. This is commonly proved by using either Euclid's lemma, the fundamental theorem of arithmetic, or the Euclidean algorithm. This is the meaning of "greatest" that is used for the generalizations of the concept of GCD.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		n, err := cmd.Flags().GetInt("n")
		if err != nil {
			return err
		}

		a, err := cmd.Flags().GetInt("a")
		if err != nil {
			return err
		}

		var s1, s2, s3, t1, t2, t3 int
		var output1, output2 string
		for i := 0; i < n; i++ {
			if n%a == 0 {
				break
			}
			if i == 0 {
				t2 = n / a
				s2 = 1
				aTmp := a
				nTmp := n
				n = a
				a = nTmp % aTmp

				output1 += pterm.Sprintln(nTmp, "=", t2, "*", aTmp, "+", a)
				output2 += pterm.Sprintln(a, "=", s2, "* n -", t2, "* a;")
			} else if i == 1 {
				t1 = t2
				s1 = s2
				multiplicand := n / a
				s2 = multiplicand * s1
				t2 = multiplicand*t1 + 1
				aTmp := a
				nTmp := n
				n = a
				a = nTmp % aTmp

				output1 += pterm.Sprintln(nTmp, "=", multiplicand, "*", aTmp, "+", a)
				output2 += pterm.Sprintln(a, "= -", s2, "* n +", t2, "* a;")
			} else {
				multiplicand := n / a
				s3 = multiplicand*s2 + s1
				t3 = multiplicand*t2 + t1
				aTmp := a
				nTmp := n
				n = a
				a = nTmp % aTmp

				if i%2 == 1 {
					output1 += pterm.Sprintln(nTmp, "=", multiplicand, "*", aTmp, "+", a)
					output2 += pterm.Sprintln(a, "= -", s3, "* n +", t3, "* a;")
				} else {
					output1 += pterm.Sprintln(nTmp, "=", multiplicand, "*", aTmp, "+", a)
					output2 += pterm.Sprintln(a, "=", s3, "* n -", t3, "* a;")
				}

				t1 = t2
				t2 = t3
				s1 = s2
				s2 = s3
			}
		}

		panels := pterm.Panels{
			{{output1}, {output2}},
		}
		_ = pterm.DefaultPanel.WithPanels(panels).WithPadding(5).Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(gcdCmd)

	gcdCmd.Flags().IntP("n", "n", 0, "n")
	gcdCmd.MarkFlagRequired("n")
	gcdCmd.Flags().IntP("a", "a", 0, "a")
	gcdCmd.MarkFlagRequired("a")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gcdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gcdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
