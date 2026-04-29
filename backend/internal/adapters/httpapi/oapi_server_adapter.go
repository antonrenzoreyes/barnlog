package httpapi

import (
	"net/http"

	openapicontract "barnlog/backend/internal/contracts/openapi"
)

var _ openapicontract.ServerInterface = (*oapiServerAdapter)(nil)

type oapiServerAdapter struct {
	system handlers
	animal animalHandlers
	upload uploadHandlers
}

func (a oapiServerAdapter) PostAnimals(w http.ResponseWriter, r *http.Request, _ openapicontract.PostAnimalsParams) {
	a.animal.createAnimal(w, r)
}

func (a oapiServerAdapter) GetHealthz(w http.ResponseWriter, r *http.Request) {
	a.system.healthz(w, r)
}

func (a oapiServerAdapter) GetReadyz(w http.ResponseWriter, r *http.Request) {
	a.system.readyz(w, r)
}

func (a oapiServerAdapter) PostUploadsAnimalPhotos(w http.ResponseWriter, r *http.Request) {
	a.upload.uploadAnimalPhoto(w, r)
}
