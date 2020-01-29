package control

import (
	"CidadesDigitaisV2/api/config"
	"CidadesDigitaisV2/api/models"
	"CidadesDigitaisV2/api/responses"
	"CidadesDigitaisV2/api/validation"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*  =========================
	FUNCAO LISTAR CD_ITENS POR ID
=========================  */

func (server *Server) GetCDItensByID(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13022)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna a rota das variaveis
	vars := mux.Vars(r)

	//	cdCodIbge armazena a chave primaria da tabela cd_itens
	cdCodIbge, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	cdCodItem armazena a chave primaria da tabela cd_itens
	cdCodItem, err := strconv.ParseUint(vars["cod_item"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	//	cdCodTipoItem armazena a chave primaria da tabela cd_itens
	cdCodTipoItem, err := strconv.ParseUint(vars["cod_tipo_item"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't parse the variable, %v\n", err))
		return
	}

	cdItens := models.CDItens{}

	//	cdItensGotten recebe o dado buscado no banco de dados
	cdItensGotten, err := cdItens.FindCDItensByID(server.DB, uint64(cdCodIbge), uint64(cdCodItem), uint64(cdCodTipoItem))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("[FATAL] It couldn't find by ID, %v\n", err))
		return
	}

	//retorna um JSON indicando que funcionou corretamente
	responses.JSON(w, http.StatusOK, cdItensGotten)

}

/*  =========================
	FUNCAO LISTAR TODOS CD_ITENS
=========================  */

func (server *Server) GetAllCDItens(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13022)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	cdItens := models.CDItens{}

	//	allCDItens armazena os dados buscados no banco de dados
	allCDItens, err := cdItens.FindAllCDItens(server.DB)
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't find in database, %v\n", formattedError))
		return
	}
	//	Retorna o Status 200 e o JSON da struct buscada
	responses.JSON(w, http.StatusOK, allCDItens)
}

/*  =========================
	FUNCAO EDITAR CD_ITENS
=========================  */

func (server *Server) UpdateCDItens(w http.ResponseWriter, r *http.Request) {

	//	Autorizacao de Modulo
	err := config.AuthMod(w, r, 13023)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, fmt.Errorf("[FATAL] Unauthorized"))
		return
	}
	//	Vars retorna as variaveis de rota
	vars := mux.Vars(r)

	//	cdCodIbge armazena a chave primaria da tabela cd_itens
	cdCodIbge, err := strconv.ParseUint(vars["cod_ibge"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//	cdCodItem armazena a chave primaria da tabela cd_itens
	cdCodItem, err := strconv.ParseUint(vars["cod_item"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//	cdCodTipoItem armazena a chave primaria da tabela cd_itens
	cdCodTipoItem, err := strconv.ParseUint(vars["cod_tipo_item"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] it couldn't read the 'body', %v\n", err))
		return
	}

	cdItens := models.CDItens{}

	//cd_itens.Prepare()

	err = json.Unmarshal(body, &cdItens)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("[FATAL] ERROR: 422, %v\n", err))
		return
	}

	if err = validation.Validator.Struct(cdItens); err != nil {
		log.Printf("[WARN] invalid information, because, %v\n", fmt.Errorf("[FATAL] validation error!, %v\n", err))
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	//	updateCDItens recebe a nova cd_itens, a que foi alterada
	updateCDItens, err := cdItens.UpdateCDItens(server.DB, uint64(cdCodIbge), uint64(cdCodItem), uint64(cdCodTipoItem))
	if err != nil {
		formattedError := config.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("[FATAL] it couldn't update in database , %v\n", formattedError))
		return
	}

	//	Retorna o Status 200 e o JSON da struct alterada
	responses.JSON(w, http.StatusOK, updateCDItens)
}
