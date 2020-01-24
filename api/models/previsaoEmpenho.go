package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

/*  =========================
	FUNCAO SALVAR PREVISAO
	EMPENHO	NO BANCO DE DADOS
=========================  */

func (previsaoEmpenho *PrevisaoEmpenho) SavePrevisaoEmpenho(db *gorm.DB) (*PrevisaoEmpenho, error) {

	//	Adiciona um novo elemento ao banco de dados
	err := db.Debug().Create(&previsaoEmpenho).Error
	if err != nil {
		return &PrevisaoEmpenho{}, err
	}
	return previsaoEmpenho, nil

}

/*  =========================
	FUNCAO LISTAR
	PREVISAO EMPENHO POR ID
=========================  */

func (previsaoEmpenho *PrevisaoEmpenho) FindPrevisaoEmpenhoByID(db *gorm.DB, previsaoEmpenhoID uint64) (*PrevisaoEmpenho, error) {

	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(PrevisaoEmpenho{}).Where("cod_previsao_empenho = ?", previsaoEmpenhoID).Take(&previsaoEmpenho).Error

	if err != nil {
		return &PrevisaoEmpenho{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &PrevisaoEmpenho{}, errors.New("Previsao_empenho Not Found")
	}

	return previsaoEmpenho, err
}

/*  =========================
	FUNCAO LISTAR TODAS PREVISAO EMPENHO
=========================  */

func (previsaoEmpenho *PrevisaoEmpenho) FindAllPrevisaoEmpenho(db *gorm.DB) (*[]PrevisaoEmpenho, error) {

	previsaoEmpenhos := []PrevisaoEmpenho{}

	// Busca todos elementos contidos no banco de dados
	err := db.Debug().Model(&PrevisaoEmpenho{}).Find(&previsaoEmpenhos).Error
	if err != nil {
		return &[]PrevisaoEmpenho{}, err
	}
	return &previsaoEmpenhos, err
}

/*  =========================
	FUNCAO EDITAR PREVISAO EMPENHO
=========================  */

func (previsaoEmpenho *PrevisaoEmpenho) UpdatePrevisaoEmpenho(db *gorm.DB, previsaoEmpenhoID uint64) (*PrevisaoEmpenho, error) {

	//	Permite a atualizacao dos campos indicados
	db = db.Debug().Model(&PrevisaoEmpenho{}).Where("cod_previsao_empenho = ?", previsaoEmpenhoID).Take(&previsaoEmpenho).UpdateColumns(
		map[string]interface{}{
			"data":           previsaoEmpenho.Data,
			"tipo":           previsaoEmpenho.Tipo,
			"ano_referencia": previsaoEmpenho.Ano_referencia,
		},
	)

	if db.Error != nil {
		return &PrevisaoEmpenho{}, db.Error
	}
	//	Busca um elemento no banco de dados a partir de sua chave primaria
	err := db.Debug().Model(&PrevisaoEmpenho{}).Where("cod_previsao_empenho = ?", previsaoEmpenhoID).Take(&previsaoEmpenho).Error
	if err != nil {
		return &PrevisaoEmpenho{}, err
	}

	// retorna o elemento que foi alterado
	return previsaoEmpenho, err
}