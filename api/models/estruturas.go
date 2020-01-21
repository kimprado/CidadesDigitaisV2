package models

//	Estrutura referente a tabela Entidade

type Entidade struct {
	Cnpj           uint64 `gorm:"primary_key;not null;size:14" json:"cnpj"`
	Nome           string `gorm:"size:50;default:null" json:"nome" validate: "alphanum":`
	Endereco       string `gorm:"size:100;default:null" json:"endereco" validate: "alphanum":`
	Numero         string `gorm:"size:10;default:null" json:"numero" validate: "alphanum":`
	Bairro         string `gorm:"size:100;default:null" json:"bairro" validate: "alphanum":`
	Cep            string `gorm:"size:8;default:null" json:"cep" validate: "alphanum":`
	Nome_municipio string `gorm:"size:50;default:null" json:"nome_municipio" validate: "alphanum":`
	Uf             string `gorm:"size:2;default:null" json:"uf" validate: "alphanum":`
	Observacao     string `gorm:"size:1000;default:null" json:"observacao" validate: "alphanum":`
}

type Lote struct {
	Cod_lote      uint32 `gorm:"primary_key;not null;size:11" json:"cod_lote" validate: "required":`
	Cnpj          string `gorm:"foreing_key:Cnpj;not null;size:14" json:"cnpj" validate: "required":`
	Contrato      string `gorm:"size:10;default:null" json:"contrato" validate: "alphanum":`
	Dt_inicio_vig string `gorm:"size:10;default:null" json:"dt_inicio_vig" validate: "alphanum":`
	Dt_final_vig  string `gorm:"size:10;default:null" json:"dt_final_vig" validate: "alphanum":`
	Dt_reajuste   string `gorm:"size:10;default:null" json:"dt_reajuste" validate: "alphanum":`
}

type Reajuste struct {
	Ano_ref    uint32  `gorm:"primary_key;not null;size:11" json:"ano_ref" validate: "required":`
	Cod_lote   uint32  `gorm:"primary key;foreing_key:Cod_lote;not null;size:11" json:"cod_lote" validate: "required":`
	Percentual float64 `gorm:"default:null" json:"percentual" validate: "alphanum":`
}

type Cd struct {
	Cod_ibge uint32 `gorm:"primary_key;foreing_key:Cod_ibge;not null;size:7" json:"cod_ibge" validate: "required":`
	Cod_lote uint32 `gorm:"foreing_key:Cod_lote;not null;size:11" json:"cod_lote" validate: "required":`
	Os_pe    string `gorm:"size:10;default:null" json:"os_pe" validate: "alphanum":`
	Data_pe  string `gorm:"default:null" json:"data_pe" validate: "alphanum":`
	Os_imp   string `gorm:"size:10;default:null" json:"os_imp" validate: "alphanum":`
	Data_imp string `gorm:"default:null" json:"data_imp" validate: "alphanum":`
}

type Cd_itens struct {
	Cod_ibge                     uint32 `gorm:"primary_key;foreing_key:Cod_ibge;not null;size:7" json:"cod_ibge" validate: "required":`
	Cod_item                     uint32 `gorm:"primary_key;foreing_key:Cod_item;not null;size:11" json:"cod_item" validate: "required":`
	Cod_tipo_item                uint32 `gorm:"primary_key;foreing_key:Cod_tipo_item;not null;size:11" json:"cod_tipo_item" validate: "required":`
	Quantidade_previsto          uint32 `gorm:"default:null;size:11" json:"quantidade_previsto" validate: "required":`
	Quantidade_projeto_executivo uint32 `gorm:"default:null;size:11" json:"quantidade_projeto_executivo" validate: "alphanum":`
	Quantidade_termo_instalacao  uint32 `gorm:"default:null;size:11" json:"quantidade_termo_instalacao" validate: "alphanum":`
}

type Lote_itens struct {
	Cod_lote      uint32  `gorm:"primary_key;foreing_key:Cod_lote;not null;size:11" json:"cod_lote" validate: "required":`
	Cod_item      uint32  `gorm:"primary_key;foreing_key:Cod_item;not null;size:11" json:"cod_item" validate: "required":`
	Cod_tipo_item uint32  `gorm:"primary_key;foreing_key:Cod_tipo_item;not null;size:11" json:"cod_tipo_item" validate: "required":`
	Preco         float64 `gorm:"default:null;size:12" json:"preco" validate: "alphanum":`
}

type Previsao_empenho struct {
	Cod_previsao_empenho uint32 `gorm:"AUTO_INCREMENT;primary_key;foreing_key:Cod_previsao_empenho;not null;size:11" json:"cod_previsao_empenho" validate: "required":`
	Cod_lote             uint32 `gorm:"foreing_key:Cod_lote;not null;size:11" json:"cod_lote" validate: "required":`
	Cod_natureza_despesa uint32 `gorm:"foreing_key:Cod_natureza_despesa;not null;size:11" json:"cod_natureza_despesa" validate: "required":`
	Data                 string `gorm:"default:null" json:"data" validate: "required":`
	Tipo                 string `gorm:"default:null;size:1" json:"tipo" validate: "alphanum":`
	Ano_referencia       uint32 `gorm:"default:null;size:11" json:"ano_referencia" validate: "alphanum":`
}

type Itens_previsao_empenho struct {
	Cod_previsao_empenho uint32  `gorm:"primary_key;;not null;size:11" json:"Cod_previsao_empenho" validate: "required":`
	Cod_item             uint32  `gorm:"primary_key;foreing_key:Cod_item;not null;size:11" json:"cod_item" validate: "required":`
	Cod_tipo_item        uint32  `gorm:"primary_key;foreing_key:Cod_tipo_item;not null;size:11" json:"cod_tipo_item" validate: "required":`
	Cod_lote             uint32  `gorm:"foreing_key:Cod_lote;not null;size:11" json:"cod_lote" validate: "required":`
	Valor                float64 `gorm:"default:null;size:12" json:"valor" validate: "alphanum":`
	Quantidade           uint32  `gorm:"default:null;size:11" json:"quantidade" validate: "alphanum":`
}