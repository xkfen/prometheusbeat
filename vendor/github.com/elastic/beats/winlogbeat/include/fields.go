// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("winlogbeat/fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzMWt1y27gVvvdTnMlNkxlJY8lJuuuLTtNk2/V0092pt02vIkHgkYSGBLj4kaw+fecAIAhStCTH2nZ9ZZHgdz4A5x+4GsMX3N/CEpm9ArDClngLfwq/CjRci9oKJW/hD1cAAO+VtExIA1xVlZL+O1gJLAsDbMtEyZYlgpDAyhJwi9KC3ddoJlcQh91eeaAxSFZhEDyhf/3TQZn09/MG/QegVmA36BmCQVkIufYPSrWGCo1hazQTuMtG+c+ESVAGLRGk91zJlVg7zUgcrESJI3pOL5mFLSsdfQnOYOExhaWfUtkczH8CG2VslBTH/6y8qA6PEb3zjxb0c5FwlJ/x47wmh4vWSDy9cIkbM6DROi2xgOXei1I1khi5BrM3FitQEnYbwTct8WzttJNSyPUAGysq/I+SZ7BpRv6abLaojVDyNJk4sFErr85+89coiQoWYDfCBFWedFX3xR9pKsayqn4RQUnXb6FgtlkHjb84obG4Batd83CldMVsZxw+sKom03vn1s5YmL21G5hdT9+OYDq7vXlz++ZmcnMzO291PSXYBUXGaIZkIBq50gXsmGnn15uUZWtzXMo7vRRWM733Y8NqcUauwOt7jTpsFJOF/2E1k4Zx2+5HWKee4OAdOuuolv9G3tha+DEPb77gfqd0cZxo8lXOoG5tihxUENZjgFor3SGw1srVx4V8Rx81HpAHiaS/rCgEjWUlCLlSZNmcGe+/vBwzaZQhesUGsGETnVl63nCy+GCzh4/QaqlFnMmBAK6KQ/RSyfVT0AnkEJqwDqC7e3YWelCTGKJ4qVzRxqj39BNqrbaiQJqmZQWzbDhsfYxvYaVVFZDSp4b2qnVBrCjmfsC8gaSRHI1R+tEoRkMn/qtJA9s3bOQnrPdvWXjrMpzAT8oYQYrrY5IBppEAR7DmOAKloRBrYVmpODI5eZSbkMYyyXEuTpjOXRwIdx8aShREoGJ8I2TfdIcknI5MSUYe18+TEgfMMz1L62xnkwoL4arj0j8GCK9iTxMe0xxRCrufZyEvMXBmjMzY8ZSfcKQZEPiIKNpoJ0ygI0wb5o6onPeNaVcTlfhm/HC+6sVPiMtflFqXGCztceka1ydD7d/9mFPzi4ZeKP7F20+09A/N7wHw8A6MZZbcb1kip5jtzTy8I5s1G6XtPESAW1ix0tCmMck3SjfyxsnKr7pOuZlyogWD8eExPx5jAuqJKJ7nE/8hxS8OW0AQxZBXT+KqofDxJIm5Xni4JjuNBCiRWDpRWlDyGJXMGXwlk/dJJmEdk1WyJZbmQFonl4Dj+cQJLnd+JYKcpLSkzK3Kfh9+DYDcUTKQKSpFuQPX0+omPT+pmVH20/Ty+XvyfSwrDnfjQpoeHMSAkjPNN8Iit05fYA4dOHiJk/UEHr55O3/7egRMVyOoaz6CStTm1SEVZSZ1ySyl9M9j8uM9NECRA0dplRmBWzpp3Qh2QhZq9wiJbsXz9RwizqCMFatEuX+2iAATJ6mx2DA7ggKXgskRrDTi0hTHZivqAwqdR0ek/yCMJYd299OYFYVGY9AcCqgYf94kGzEbposd09gKG4EzjpXlHj6+e59zaPzIF7dELdGiab3JX/NnA2Lb9ykN7ua0LSjkvuR4WGw/OumAOqThSW6oVsUFwkO2ArUqgm8bFOWe65p6kghv0LWamvHLTapFPBRGFdhFV5AQH1nCc4PreYICGlSsPpTEpFTW978uJi6DHJZ5yYQlk8s7ucsxsRdI2QblBtymjvad29a7vHgfWrmfhCzVeonMvjiv9/vEti8MNwgGO2RHGlxtN+vdT3ehfHMGC7AKNLLC+7vQ55r48fVh6ZzwXuyE9HilWr9IudinEHngOy/phyhJaXjRjF0L2Y5PaGk8vaZvJlcd5gfvff5coBFrGVPBRvQ96i1qmF1f3ySI7PXs+vr6oFdqJnAn05B/CmPZqNv/I7kJTsiVZsZqF9Kene/GNmQm8GMPisSXzKJu5SaoKH90ZPVSV7zVMmDOqopZwSkeXrW6TZmYiYWi32Plu+a0PrTDNOM0JYqcjWI3u9Oq9ne9zucJvW47RCxNo988fUyjGbdiK+y+LcKHjLdV9ibaHtH2d7Au1dInCy4kxKJAacVKoA5t1/TbhMLMaU1sGy7eABKcn0ns1zKNULtlKcwGC9gJG0vyTIAfwrRt6j+Tu8QkodsV4KqqnUWdd3+OL8N5Np93iBoZh2361u4/bVCCM6QpcR+7q0C6tGOaVGkUJu4lcCahEKsV6qAF/pymc9ay6LeMCW0e+47w5Mb1k9QhOb+xqZGLleBeYWmThQk6SWZWORsSTHzgpTNii35/E8zCGdSe8sKfae2V83vNWW2dbo3LG4OSTUFAmaNI7XEgoxxwNTXTrEKL2lAwWLTLs/AyaNEKWPhR08Wo5eSfzBYjf2hgFCg5giVyRnbfWmCGTmhOBjzRVj3IdEnKm0irVUNzcOd6xpp1v78uLLX2E4JPOkJMW2ZVsCblNE86PXS6FDXFXMqbEJ0GMyxfjJq8ZMaI1R6YHCRSqvWvac/t9nqby/qD2WkV+f2o6M2StnFRJiyfJme2GzaZwsSi6bk8dsranzRusbzk0nvA7n7Tc41+L1Zkpn6IV9ncVYejtYRFYZzE3jtOJdnIN5R05ecwgk9MS+/U/PnJCN65Qth2LFmXf5Tg/sxE6fRhtzs/ezo4d3ry5CNeuzOHUXVI/vzwRO65/jMcK9G6Ku7DZUEaV2LKLGiJwoLa5mpBzr71f+HOQY94mM1cumqJ5xE/z2CiKQTcQ8vJ892V0MYm29HCWpS5126s3H/YWlmEngYlUXaDOr4K3iK8x7YKMfiLI2dHsSZduOiCaWR8E3OTij2IylXRfF/OPt/MPiesJvM+zJCJzOzz29efj2fnr0adjZH4YHtcdqIsYYlwfbBjlNUW899C/tZwoG3yfjBBUSmnVel9oT9PX6HXXasmUT/8LBrHsVOuLMJNlg3bonenbeKf5XhertCwyKa/yKNYd7FUnR0aX8QgAyIUuBLSh/KOf2TmS1DHMIr00O7rUDA08as/r300Tc5yX8/quhTxUYwz5HtbM9ox0/Oxvbk3h7Hztfs/qkg67/YvA9GO6g9F8Ni+OyvbOXPfWsgBhvRmKD0f4hYSoUsnGAPpVX4B5mVfJZQGg3orOHrePU3IXNurg+sy5ssl7YHwfpvW4OlxZnGtdBQXu8uHXvulklBr7FYIbcLWb1y8ojgUQ1VIkJuCpr/aGwrTl9TjhNhX4/DiXC1ONdX/sAyE55WBWW02MJlJ64eerN7p9sH9eDp+M55NxzdvXk9f31x/O/tmPLt+M/39dDqbXo+nN99Ob755ffP22/H0+np6etqNOhnkTlOczJzly/u7D68ag2ecKyctMGMUF37nOpPvXOlLT+9WeXeK+3YxaDSq3AbbuL/74DOoeJXTNhfsfEnte8Z5TesfFqpiQsbCNjyihVw0zdMmLVEVaX/Ry5AnOTcq1AthuNqizom2LMmk7u8+mBFo3ArcpYuxq6xw56ELakKSwWzjtJclVlCxPSz71UCa3SUdXV4CPrJdj9+8zJb2kpwCYrD583iRpTGoMKblQzSzVvdlIkS8sPTkReueEH+N68wvzA6UIr8zIXKJgXr6oTpWTT+/wtRsB//6+ANorDUalDbG9TwBUEt/HBKNrGkLeb/ZViKN/1Sy3GdHG4TVD1tgXF0rnerEfqux2wF/+VFwrYxa2V5jnTyDxB3qV30fnl80F5KXLh7pFrhirgw3bCtnLDkQlMTTvzYYatdF/Gb+UJUhcWrvmzPS2+4183gFneosRplCIbaicKxs06fcWdKin1rw0PRfuTIUb1q5ZYlmo1Tn9KB2ulYGTcgwfNsxpinRR2qkVQ5z6zpomjhXVc3ywjR22HIgImoVFIKtpTLJ3ZnJ1X8DAAD//xTx/qA="
}
