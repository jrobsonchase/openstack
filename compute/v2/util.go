package v2

import (
	"errors"
)

var (
	ErrServerNotFound error = errors.New("Server not found.")
	ErrImageNotFound error = errors.New("Image not found.")
	ErrFlavorNotFound error = errors.New("Flavor not found.")
	ErrNetworkNotFound error = errors.New("Flavor not found.")
	ErrMissingName = errors.New("Missing Name.")
	ErrMissingFlavor = errors.New("Missing Flavor.")
	ErrMissingImage = errors.New("Missing Image.")
	ErrMissingNetworks = errors.New("Missing Networks.")
)

func (cc *ComputeClient) ServerByName(name string) (*Server, error) {
	srvs,err := cc.Servers()
	if err != nil {
		return nil,err
	}
	for _,v := range srvs {
		if v.Name == name {
			srv,err := cc.Details(v)
			if err == nil {
				return srv,nil
			}
			return nil,err
		}
	}
	return nil,ErrServerNotFound
}

func (cc *ComputeClient) ImageByName(name string) (*Image,error) {
	imgs,err := cc.Images()
	if err != nil {
		return nil,err
	}
	for _,v := range imgs {
		if v.Name == name {
			img,err := cc.ImageDetails(v)
			if err == nil {
				return img,nil
			}
			return nil,err
		}
	}
	return nil,ErrImageNotFound
}

func (cc *ComputeClient) FlavorByName(name string) (*Flavor,error) {
	flavs,err := cc.Flavors()
	if err != nil {
		return nil,err
	}
	for _,v := range flavs {
		if v.Name == name {
			flav,err := cc.FlavorDetails(v)
			if err == nil {
				return flav,nil
			}
			return nil,err
		}
	}
	return nil,ErrFlavorNotFound
}

func (cc *ComputeClient) NetworkByName(name string) (*Network,error) {
	nets,err := cc.Networks()
	if err != nil {
		return nil,err
	}
	for _,v := range nets {
		if v.Label == name {
			return v,nil
		}
	}
	return nil,ErrNetworkNotFound
}

func (cc *ComputeClient) Create(srv *Server) (*Server,error) {
	if srv.Name == "" {
		return nil,ErrMissingName
	}
	if srv.FlavorRef == "" {
		if srv.Flavor.Name != "" {
			flav,err := cc.FlavorByName(srv.Flavor.Name)
			if err != nil {
				return  nil,err
			}
			srv.FlavorRef = flav.Id
		} else {
			return nil,ErrMissingFlavor
		}
	}
	if srv.ImageRef == "" {
		if srv.Image.Name != "" {
			img,err := cc.ImageByName(srv.Image.Name)
			if err != nil {
				return  nil,err
			}
			srv.ImageRef = img.Id
		} else {
			return nil,ErrMissingImage
		}
	}
	if srv.Networks == nil {
		if srv.NetNames != nil {
			srv.Networks = make([]NetConf,len(srv.NetNames))
			i := 0
			for _,v := range srv.NetNames {
				net,err := cc.NetworkByName(v)
				switch err {
				case ErrNetworkNotFound:
					continue
				case nil:
					srv.Networks[i] = NetConf{Uuid: net.Id}
					i++
				default:
					return nil,err
				}
			}
			srv.Networks = srv.Networks[:i]
		} else {
			return nil,ErrMissingNetworks
		}
	}
	return cc.create(srv)
}

