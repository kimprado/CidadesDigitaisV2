package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*	=========================
	STRUCT CATEGORIA
=========================	*/

type Categoria struct {
	CodCategoria uint32 `gorm:"primary_key;auto_increment;not null" json:"cod_categoria"`
	Descricao    string `gorm:"default:null" json:"descricao"`
}

/*  =========================
	FUNCAO SALVAR CATEGORIA NO BANCO DE DADOS
=========================  */

func (categoria *Categoria) SaveCategoria(db *gorm.DB) (*Categoria, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&categoria).Error
	if err != nil {
		return &Categoria{}, err
	}

	return categoria, nil
}

/*  =========================
	FUNCAO LISTAR CATEGORIA POR ID
=========================  */

func (categoria *Categoria) FindCategoriaByID(db *gorm.DB, codCategoria uint32) (*Categoria, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(Categoria{}).Where("cod_categoria = ?", codCategoria).Take(&categoria).Error
	if err != nil {
		return &Categoria{}, err
	}

	return categoria, err
}

/*  =========================
	FUNCAO LISTAR TODAS CATEGORIA
=========================  */

func (categoria *Categoria) FindAllCategoria(db *gorm.DB) (*[]Categoria, error) {

	allCategoria := []Categoria{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&Categoria{}).Find(&allCategoria).Error
	if err != nil {
		return &[]Categoria{}, err
	}

	return &allCategoria, err
}

/*  =========================
	FUNCAO EDITAR CATEGORIA
=========================  */

func (categoria *Categoria) UpdateCategoria(db *gorm.DB, codCategoria uint32) (*Categoria, error) {

	//	Permite a atualizacao dos campos indicados
	err := db.Debug().Exec("UPDATE categoria SET descricao = ? WHERE cod_categoria = ?", categoria.Descricao, codCategoria).Error
	if err != nil {
		return &Categoria{}, err
	}

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err = db.Debug().Model(&Categoria{}).Where("cod_categoria = ?", codCategoria).Take(&categoria).Error
	if err != nil {
		return &Categoria{}, err
	}

	return categoria, err
}

/*  =========================
	FUNCAO DELETAR CATEGORIA POR ID
=========================  */

func (categoria *Categoria) DeleteCategoria(db *gorm.DB, codCategoria uint32) error {

	//	Deleta um elemento contido no banco de dados a partir de sua chave primaria
	db = db.Debug().Model(&Categoria{}).Where("cod_categoria = ?", codCategoria).Take(&Categoria{}).Delete(&Categoria{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return errors.New("Categoria not found")
		}
		return db.Error
	}

	return db.Error
}
