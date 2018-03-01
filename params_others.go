// +build !linux,!darwin,!windows,!openbsd

package water

// PlatformSpeficParams
type PlatformSpecificParams struct {
}

func defaultPlatformSpecificParams() PlatformSpecificParams {
	return PlatformSpecificParams{}
}
