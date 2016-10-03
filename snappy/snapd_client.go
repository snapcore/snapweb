/*
 * Copyright (C) 2016 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package snappy

import (
	"log"
	"time"

	"github.com/snapcore/snapd/client"
	"gopkg.in/ini.v1"
)

var timesyncdConfigurationFilePath = "/etc/systemd/timesyncd.conf"

type DeviceBranding struct {
	Name      string `json:"name"`
	Subname   string `json:"subname"`
	Icon      string `json:"icon"`
	IconSvg   string `json:"iconsvg"`
	Link      string `json:"link"`
	LinkSrc   string `json:"linksrc"`
	TextColor string `json:"textcolor"`
	LinkColor string `json:"linkcolor"`
}

// SnapdClient is a client of the snapd REST API
type SnapdClient interface {
	Icon(name string) (*client.Icon, error)
	Snap(name string) (*client.Snap, *client.ResultInfo, error)
	List(names []string) ([]*client.Snap, error)
	Find(opts *client.FindOptions) ([]*client.Snap, *client.ResultInfo, error)
	Install(name string, options *client.SnapOptions) (string, error)
	Remove(name string, options *client.SnapOptions) (string, error)
	ServerVersion() (*client.ServerVersion, error)
}

// ClientAdapter adapts our expectations to the snapd client API.
type ClientAdapter struct {
	snapdClient *client.Client
}

// NewClientAdapter creates a new ClientAdapter for use in snapweb.
func NewClientAdapter() *ClientAdapter {
	return &ClientAdapter{
		snapdClient: client.New(nil),
	}
}

// Icon returns the Icon belonging to an installed snap.
func (a *ClientAdapter) Icon(name string) (*client.Icon, error) {
	return a.snapdClient.Icon(name)
}

// Snap returns the most recently published revision of the snap with the
// provided name.
func (a *ClientAdapter) Snap(name string) (*client.Snap, *client.ResultInfo, error) {
	return a.snapdClient.Snap(name)
}

// List returns the list of all snaps installed on the system
// with names in the given list; if the list is empty, all snaps.
func (a *ClientAdapter) List(names []string) ([]*client.Snap, error) {
	return a.snapdClient.List(names)
}

// Find returns a list of snaps available for install from the
// store for this system and that match the query
func (a *ClientAdapter) Find(opts *client.FindOptions) ([]*client.Snap, *client.ResultInfo, error) {
	return a.snapdClient.Find(opts)
}

// Install adds the snap with the given name from the given channel (or
// the system default channel if not).
func (a *ClientAdapter) Install(name string, options *client.SnapOptions) (string, error) {
	return a.snapdClient.Install(name, options)
}

// Remove removes the snap with the given name.
func (a *ClientAdapter) Remove(name string, options *client.SnapOptions) (string, error) {
	return a.snapdClient.Remove(name, options)
}

// ServerVersion returns information about the snapd server.
func (a *ClientAdapter) ServerVersion() (*client.ServerVersion, error) {
	return a.snapdClient.ServerVersion()
}

// internal
func readNTPServer() string {
	timesyncd, err := ini.Load(timesyncdConfigurationFilePath)
	if err != nil {
		log.Println("readNTPServer: unable to read ",
			timesyncdConfigurationFilePath)
		return ""
	}

	section, err := timesyncd.GetSection("Time")
	if err != nil || !section.HasKey("NTP") {
		log.Println("readNTPServer: no NTP servers are set ",
			timesyncdConfigurationFilePath)
		return ""
	}

	return section.Key("NTP").Strings(" ")[0]
}

// GetCoreConfig gets some aspect of core configuration
// XXX: current assumption, asking for timezone info
func GetCoreConfig(keys []string) (map[string]interface{}, error) {
	var dt = time.Now()
	_, offset := dt.Zone()

	return map[string]interface{}{
		"Date":      dt.Format("2006-01-02"), // Format for picker
		"Time":      dt.Format("15:04"),      // Format for picker
		"Timezone":  float64(offset) / 60 / 60,
		"NTPServer": readNTPServer(),
	}, nil
}

func getDefaultBranding() snappy.DeviceBranding {
	return branding{
		Name:      "Ubuntu",
		Subname:   "",
		Link:      "https://ubuntu.com",
		LinkSrc:   "https://ubuntu.com",
		TextColor: "",
		LinkColor: "",
		Icon:      "",
		IconSvg:   "<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"107\" height=\"25\" viewBox=\"0 0 107 25\"><g><circle fill=\"#e95420\" cx=\"100.596\" cy=\"7.374\" r=\"6.403\"/><path fill=\"#fff\" d=\"M96.6 6.605c-.424 0-.768.344-.768.77 0 .423.344.768.768.768.425 0 .772-.345.772-.768 0-.426-.347-.77-.772-.77zm5.494 3.498c-.365.21-.492.682-.282 1.047.214.372.685.497 1.05.284.37-.21.497-.682.282-1.05-.212-.37-.682-.495-1.05-.28zm-3.822-2.728c0-.76.378-1.434.956-1.84l-.562-.943c-.675.45-1.176 1.138-1.384 1.943.246.2.398.5.398.84 0 .337-.152.64-.398.838.208.805.71 1.492 1.384 1.94l.562-.94c-.578-.406-.956-1.08-.956-1.838zm2.246-2.25c1.177 0 2.14.902 2.24 2.052l1.098-.02c-.055-.845-.425-1.61-.994-2.165-.295.108-.634.093-.925-.078-.292-.168-.473-.452-.524-.76-.285-.078-.584-.12-.895-.12-.53 0-1.032.124-1.48.344l.535.958c.287-.135.61-.21.945-.21zm0 4.497c-.336 0-.658-.076-.945-.21l-.535.958c.447.22.95.344 1.48.344.312 0 .61-.04.895-.12.05-.31.232-.595.524-.762.294-.168.63-.186.925-.073.57-.563.938-1.324.994-2.17l-1.098-.016c-.1 1.147-1.063 2.048-2.24 2.048zm1.576-4.976c.368.213.838.088 1.05-.282.215-.367.088-.84-.28-1.052-.366-.21-.837-.085-1.05.283-.213.367-.086.838.28 1.05z\"/><path fill=\"#e95420\" d=\"M12.807 24.177c-.65.162-1.51.337-2.577.518-1.068.184-2.304.276-3.704.276-1.222 0-2.247-.178-3.082-.534-.833-.355-1.503-.858-2.01-1.51-.51-.65-.875-1.418-1.098-2.3C.11 19.738 0 18.76 0 17.68V8.778h2.835v8.293c0 1.933.304 3.316.916 4.148.61.834 1.637 1.25 3.08 1.25.304 0 .62-.01.945-.032.326-.02.632-.045.914-.075.284-.03.544-.06.778-.09.234-.032.4-.065.504-.108V8.778h2.835v15.4zM19.728 9.478c.347-.222.87-.458 1.572-.7.7-.243 1.507-.367 2.424-.367 1.137 0 2.148.206 3.033.61.886.406 1.632.977 2.243 1.71.607.73 1.07 1.604 1.385 2.62.317 1.017.474 2.136.474 3.356 0 1.28-.188 2.433-.566 3.46-.377 1.026-.91 1.895-1.6 2.606-.692.714-1.524 1.26-2.502 1.647-.973.387-2.07.58-3.293.58-1.322 0-2.49-.093-3.507-.276-1.016-.183-1.85-.367-2.5-.55V1.46l2.836-.488v8.506h-.002zm0 12.684c.285.084.686.16 1.206.23.517.073 1.16.107 1.935.107 1.522 0 2.745-.505 3.658-1.51.915-1.007 1.373-2.435 1.373-4.284 0-.813-.08-1.575-.244-2.288-.163-.712-.428-1.326-.793-1.845-.366-.52-.84-.923-1.416-1.22-.582-.296-1.278-.443-2.09-.443-.774 0-1.484.133-2.135.395-.65.267-1.15.54-1.495.825v10.032zM46.684 24.177c-.65.162-1.51.337-2.576.518-1.067.184-2.303.276-3.707.276-1.218 0-2.245-.178-3.078-.534-.833-.355-1.504-.858-2.01-1.51-.512-.65-.876-1.418-1.1-2.3-.223-.887-.334-1.866-.334-2.945V8.78h2.835v8.293c0 1.933.304 3.316.914 4.148.61.834 1.636 1.25 3.08 1.25.305 0 .622-.01.947-.032.325-.02.63-.045.915-.075.283-.03.543-.06.778-.09.232-.033.4-.066.503-.11V8.78h2.836v15.4h-.002zM50.77 9.236c.65-.164 1.513-.335 2.592-.52 1.077-.18 2.316-.273 3.72-.273 1.262 0 2.307.178 3.14.533.835.356 1.5.853 1.998 1.494.498.64.85 1.408 1.053 2.3.202.898.304 1.882.304 2.96v8.904H60.74V16.34c0-.977-.066-1.81-.197-2.5-.132-.694-.35-1.25-.656-1.677-.303-.43-.712-.738-1.22-.93-.507-.197-1.14-.292-1.89-.292-.305 0-.62.013-.943.033-.326.02-.637.046-.93.076-.3.03-.56.066-.795.106-.236.04-.403.072-.504.09v13.387H50.77V9.236zM70.132 8.778h6.008v2.378h-6.008v7.32c0 .79.062 1.45.185 1.965.12.52.304.928.548 1.222.244.292.55.5.914.625.366.122.793.182 1.28.182.873 0 1.564-.097 2.075-.29.507-.194.853-.33 1.035-.413l.61 2.32c-.284.142-.788.324-1.51.547-.72.225-1.538.338-2.453.338-1.078 0-1.968-.137-2.67-.413-.7-.276-1.264-.686-1.69-1.235-.428-.548-.728-1.224-.9-2.027-.174-.805-.26-1.734-.26-2.79V4.356l2.834-.49V8.78h.002zM91.57 24.177c-.653.162-1.51.337-2.58.518-1.066.184-2.3.276-3.7.276-1.222 0-2.247-.178-3.08-.534-.837-.355-1.507-.858-2.014-1.51-.51-.65-.873-1.418-1.1-2.3-.22-.887-.335-1.866-.335-2.945V8.78h2.84v8.293c0 1.933.304 3.316.913 4.148.61.834 1.637 1.25 3.08 1.25.305 0 .62-.01.946-.032.322-.02.628-.045.913-.075.284-.03.546-.06.777-.09.234-.033.403-.066.505-.11V8.78h2.837v15.4h-.002z\"/></g></svg>",
	}
}

func GetBrandingData(h snappy.Handler) (snappy.DeviceBranding, error) {
	snap := h.getSnapByType("gadget")
	if snap == nil || len(snap) == 0 {
		return getDefaultBranding()
	}
	// TODO get snapd conf
	return snappy.DeviceBranding{
		Name:      snap.Name,
		Subname:   snap.Description,
		Icon:      snap.Icon,
		Link:      "https://ubuntu.com",
		LinkSrc:   "https://ubuntu.com",
		TextColor: "",
		LinkColor: "",
	}, nil
}
