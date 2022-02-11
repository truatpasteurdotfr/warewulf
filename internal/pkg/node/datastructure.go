package node

import (
	"github.com/hpcng/warewulf/internal/pkg/util"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
)

/******
 * YAML data representations
 ******/

type nodeYaml struct {
	NodeProfiles map[string]*NodeConf
	Nodes        map[string]*NodeConf
}

type NodeConf struct {
	Comment        string              `yaml:"comment,omitempty"`
	ClusterName    string              `yaml:"cluster name,omitempty"`
	ContainerName  string              `yaml:"container name,omitempty"`
	Ipxe           string              `yaml:"ipxe template,omitempty"`
	KernelVersion  string              `yaml:"kernel version,omitempty"`
	KernelArgs     string              `yaml:"kernel args,omitempty"`
	IpmiUserName   string              `yaml:"ipmi username,omitempty"`
	IpmiPassword   string              `yaml:"ipmi password,omitempty"`
	IpmiIpaddr     string              `yaml:"ipmi ipaddr,omitempty"`
	IpmiNetmask    string              `yaml:"ipmi netmask,omitempty"`
	IpmiPort       string              `yaml:"ipmi port,omitempty"`
	IpmiGateway    string              `yaml:"ipmi gateway,omitempty"`
	IpmiInterface  string              `yaml:"ipmi interface,omitempty"`
	RuntimeOverlay string              `yaml:"runtime overlay,omitempty"`
	SystemOverlay  string              `yaml:"system overlay,omitempty"`
	Init           string              `yaml:"init,omitempty"`
	Root           string              `yaml:"root,omitempty"`
	AssetKey       string              `yaml:"asset key,omitempty"`
	Discoverable   string              `yaml:"discoverable,omitempty"`
	Profiles       []string            `yaml:"profiles,omitempty"`
	NetDevs        map[string]*NetDevs `yaml:"network devices,omitempty"`
	Tags           map[string]string   `yaml:"tags,omitempty"`
}

type NetDevs struct {
	Type    string
	OnBoot  string
	Device  string
	Hwaddr  string
	Ipaddr  string
	IpCIDR  string `yaml:"ipcidr,omitempty"`
	Prefix  string `yaml:"prefix,omitempty"`
	Netmask string
	Gateway string `yaml:"gateway"`
	Default string
}

/******
 * Internal code data representations
 ******/

type Entry struct {
	value    string
	altvalue string
	from     string
	def      string
}

type NodeInfo struct {
	Id             Entry
	Cid            Entry
	Comment        Entry
	ClusterName    Entry
	ContainerName  Entry
	Ipxe           Entry
	KernelVersion  Entry
	KernelArgs     Entry
	IpmiIpaddr     Entry
	IpmiNetmask    Entry
	IpmiPort       Entry
	IpmiGateway    Entry
	IpmiUserName   Entry
	IpmiPassword   Entry
	IpmiInterface  Entry
	RuntimeOverlay Entry
	SystemOverlay  Entry
	Root           Entry
	Discoverable   Entry
	Init           Entry //TODO: Finish adding this...
	AssetKey       Entry
	Profiles       []string
	GroupProfiles  []string
	NetDevs        map[string]*NetDevEntry
	Tags           map[string]*Entry
}

type NetDevEntry struct {
	Type    Entry
	OnBoot  Entry
	Device  Entry
	Hwaddr  Entry
	Ipaddr  Entry
	IpCIDR  Entry
	Prefix  Entry
	Netmask Entry
	Gateway Entry
	Default Entry
}

func init() {
	// Check that nodes.conf is found
	if !util.IsFile(ConfigFile) {
		wwlog.Printf(wwlog.WARN, "Missing node configuration file\n")
		// just return silently, as init is also called for bash_completion
		return
	}
}
