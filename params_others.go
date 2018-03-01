// +build !linux,!darwin,!windows,!openbsd,!freebsd

package water

// PlatformSpeficParams
type PlatformSpecificParams struct {
}

func defaultPlatformSpecificParams() PlatformSpecificParams {
	return PlatformSpecificParams{}
}
