package todo

type TodoRequest struct {
	Kegiatan  string `json:"kegiatan" form:"kegiatan" validate:"required"`
	Deskripsi string `json:"deskripsi" form:"deskripsi" validate:"required"`
	Deadline  string `json:"deadline" form:"deadline" validate:"required,datetime=2006-01-02"`
}
