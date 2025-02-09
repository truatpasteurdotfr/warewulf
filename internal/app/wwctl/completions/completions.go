package completions

import (
	"github.com/spf13/cobra"

	"github.com/warewulf/warewulf/internal/pkg/hostlist"
	"github.com/warewulf/warewulf/internal/pkg/image"
	"github.com/warewulf/warewulf/internal/pkg/kernel"
	"github.com/warewulf/warewulf/internal/pkg/node"
	"github.com/warewulf/warewulf/internal/pkg/overlay"
)

func NodeKernelVersion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var kernelVersions []string
	registry, err := node.New()
	if err != nil {
		return kernelVersions, cobra.ShellCompDirectiveNoFileComp
	}
	nodes := hostlist.Expand(args)
	for _, id := range nodes {
		if node_, err := registry.GetNode(id); err != nil {
			continue
		} else if node_.ImageName != "" {
			kernels := kernel.FindKernels(node_.ImageName)
			for _, kernel_ := range kernels {
				kernelVersions = append(kernelVersions, kernel_.Version(), kernel_.Path)
			}
		}
	}
	return kernelVersions, cobra.ShellCompDirectiveNoFileComp
}

func ProfileKernelVersion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var kernelVersions []string
	registry, err := node.New()
	if err != nil {
		return kernelVersions, cobra.ShellCompDirectiveNoFileComp
	}
	for _, id := range args {
		if profile, err := registry.GetProfile(id); err != nil {
			continue
		} else if profile.ImageName != "" {
			kernels := kernel.FindKernels(profile.ImageName)
			for _, kernel_ := range kernels {
				kernelVersions = append(kernelVersions, kernel_.Version(), kernel_.Path)
			}
		}
	}
	return kernelVersions, cobra.ShellCompDirectiveNoFileComp
}

func Images(num int) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if num <= 0 || len(args) < num {
			if sources, err := image.ListSources(); err != nil {
				return sources, cobra.ShellCompDirectiveNoFileComp
			}
		}
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
}

func Nodes(num int) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if num <= 0 || len(args) < num {
			if registry, err := node.New(); err == nil {
				return registry.ListAllNodes(), cobra.ShellCompDirectiveNoFileComp
			}
		}
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
}

func Profiles(num int) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if num <= 0 || len(args) < num {
			if registry, err := node.New(); err == nil {
				return registry.ListAllProfiles(), cobra.ShellCompDirectiveNoFileComp
			}
		}
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
}

func Overlays(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	list := overlay.FindOverlays()
	return list, cobra.ShellCompDirectiveNoFileComp
}

func OverlayFiles(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	ret, _ := overlay.OverlayGetFiles(args[0])
	return ret, cobra.ShellCompDirectiveNoFileComp
}

func OverlayAndFiles(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) == 0 {
		return Overlays(cmd, args, toComplete)
	} else {
		return OverlayFiles(cmd, args, toComplete)
	}
}

func None(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return nil, cobra.ShellCompDirectiveNoFileComp
}

func LocalFiles(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return nil, cobra.ShellCompDirectiveDefault
}
