package control

import (
	"CidadesDigitaisV2/api/config"
	"CidadesDigitaisV2/api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) CreateHandler() (r *mux.Router) {

	//	CRIA UM ROTEADOR

	r = s.Router
	//	HOME
	r.HandleFunc("/{id}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.Home))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM USUARIO
	=========================	*/

	//ROTA DE LOGIN
	r.HandleFunc("/read/usuario/login", middlewares.SetMiddleJSON(s.Login)).Methods(http.MethodPost)

	//LISTA USUARIOS
	r.HandleFunc("/read/usuario", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetUsers))).Methods(http.MethodGet)

	//SALVA USUARIO
	r.HandleFunc("/read/usuario/createuser", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateUser))).Methods(http.MethodPost)

	//LISTA USUARIO
	r.HandleFunc("/read/usuario/{id}/{modulo}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetUser))).Methods(http.MethodGet)

	//EDITA O USUARIO
	r.HandleFunc("/read/usuario/{id}/{modulo}", middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateUser))).Methods(http.MethodPut)

	//APAGA O USUARIO
	//	r.HandleFunc(config.USER_ID_PATH, middlewares.SetMiddleAuth(s.DeleteUser)).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM ENTIDADE
	=========================	*/

	//	LISTA ENTIDADE
	r.HandleFunc(config.ENTIDADE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllEntidade))).Methods(http.MethodGet)

	//	SALVA ENTIDADE
	r.HandleFunc(config.ENTIDADE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateEntidade))).Methods(http.MethodPost)

	//	EDITA ENTIDADE
	r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateEntidade))).Methods(http.MethodPut)

	//	LISTA ENTIDADE POR ID
	r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetEntidadeByID))).Methods(http.MethodGet)

	//	APAGA ENTIDADE POR ID
	r.HandleFunc(config.ENTIDADE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteEntidade))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS DE CONTATO
	=========================	*/

	//LISTA CONTATO
	r.HandleFunc(config.CONTATO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllContato))).Methods(http.MethodGet)

	//SALVA CONTATO
	r.HandleFunc(config.CONTATO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.SaveContato))).Methods(http.MethodPost)

	//EDITA CONTATO ID
	r.HandleFunc(config.CONTATO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateContato))).Methods(http.MethodPut)

	//LISTA CONTATO ID
	r.HandleFunc(config.CONTATO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetContatoByID))).Methods(http.MethodGet)

	//	APAGA LOTE POR ID
	r.HandleFunc(config.CONTATO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteContato))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS DE TELEFONE
	=========================	*/

	//LISTA TELEFONE
	r.HandleFunc(config.TELEFONE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllTelefone))).Methods(http.MethodGet)

	//SALVA TELEFONE
	r.HandleFunc(config.TELEFONE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.SaveTelefone))).Methods(http.MethodPost)

	//APAGA TELEFONE ID
	r.HandleFunc(config.TELEFONE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteTelefone))).Methods(http.MethodDelete)

	//LISTA TELEFONE ID
	r.HandleFunc(config.TELEFONE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetTelefoneByID))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM LOTE
	=========================	*/

	//	LISTA LOTE
	r.HandleFunc(config.LOTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllLote))).Methods(http.MethodGet)

	//	SALVA LOTE
	r.HandleFunc(config.LOTE_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateLote))).Methods(http.MethodPost)

	//	EDITA LOTE
	r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateLote))).Methods(http.MethodPut)

	//	LISTA LOTE POR ID
	r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetLoteByID))).Methods(http.MethodGet)

	//	APAGA LOTE POR ID
	r.HandleFunc(config.LOTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteLote))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM CD
	=========================	*/

	//	LISTA CD
	r.HandleFunc(config.CD_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllCD))).Methods(http.MethodGet)

	//	SALVA CD
	r.HandleFunc(config.CD_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateCD))).Methods(http.MethodPost)

	//	EDITA CD
	r.HandleFunc(config.CD_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateCD))).Methods(http.MethodPut)

	//	LISTA CD POR ID
	r.HandleFunc(config.CD_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetCDByID))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM CD_ITENS
	=========================	*/

	//	LISTA CD_ITENS
	r.HandleFunc(config.CD_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllCDItens))).Methods(http.MethodGet)

	//	EDITA CD_ITENS
	r.HandleFunc(config.CD_ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateCDItens))).Methods(http.MethodPut)

	//	LISTA CD_ITENS POR ID
	r.HandleFunc(config.CD_ITENS_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetCDItensByID))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM REAJUSTE
	=========================	*/

	//	LISTA REAJUSTE
	//	NAO EXITE

	//	SALVA REAJUSTE
	r.HandleFunc(config.REAJUSTE_PATH_CREATEREAJUSTE, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreateReajuste))).Methods(http.MethodPost)

	//	LISTA REAJUSTE POR ID
	r.HandleFunc(config.REAJUSTE_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetReajusteByID))).Methods(http.MethodGet)

	//	APAGA REAJUSTE (lote_cod_lote, ano_ref)
	r.HandleFunc(config.REAJUSTE_DEL, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.DeleteReajuste))).Methods(http.MethodDelete)

	/*	=========================
			ROTAS EM LOTE ITENS
	=========================	*/

	//	EDITA LOTE ITENS
	r.HandleFunc(config.LOTE_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateLoteItens))).Methods(http.MethodPut)

	//	LISTA LOTE ITENS POR ID
	r.HandleFunc(config.LOTE_ITENS_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetLoteItensByID))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM PREVISAO EMPENHO
	=========================	*/

	//	LISTA PREVISAO EMPENHO
	r.HandleFunc(config.PREVISAO_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetAllPrevisaoEmpenho))).Methods(http.MethodGet)

	//	SALVA PREVISAO EMPENHO
	r.HandleFunc(config.PREVISAO_EMPENHO_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.CreatePrevisaoEmpenho))).Methods(http.MethodPost)

	//	EDITA PREVISAO EMPENHO
	r.HandleFunc(config.PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdatePrevisaoEmpenho))).Methods(http.MethodPut)

	//	LISTA PREVISAO EMPENHO POR ID
	r.HandleFunc(config.PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetPrevisaoEmpenhoByID))).Methods(http.MethodGet)

	/*	=========================
			ROTAS EM ITENS PREVISAO EMPENHO
	=========================	*/

	//	EDITA ITENS PREVISAO EMPENHO
	r.HandleFunc(config.ITENS_PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.UpdateItensPrevisaoEmpenho))).Methods(http.MethodPut)

	//	LISTA ITENS PREVISAO EMPENDO
	r.HandleFunc(config.ITENS_PREVISAO_EMPENHO_ID_PATH, middlewares.SetMiddleJSON(middlewares.SetMiddleAuth(s.GetItensPrevisaoEmpenhoByID))).Methods(http.MethodGet)

	/*
		//**********Rotas em Assunto
		//----------Rotas de Assunto
		//lista assunto
		r.HandleFunc().Methods(http.MethodGet)

		//salva assunto
		r.HandleFunc().Methods(http.MethodPost)

		//lista assunto por ID
		r.HandleFunc().Methods(http.MethodGet)

		//edita assunto ID
		r.HandleFunc().Methods(http.MethodPut)

		//apaga assunto ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Categoria
		//Rotas de Categoria
		//lista categotia
		r.HandleFunc().Methods(http.MethodGet)

		//salva categoria
		r.HandleFunc().Methods(http.MethodPost)

		//edita categoria
		r.HandleFunc().Methods(http.MethodPut)

		//lista categoria por ID
		r.HandleFunc().Methods(http.MethodGet)

		//deleta categoria ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Cidades Digitais
			//---------Rotas de Estado---------//
			//---------Rotas de Municipios da Cidade Digital---------//
			//---------Rotas de Cidade Digital---------//
			//---------Rotas de Itens da Cidade Digital---------//
			//---------Rotas de Processo da Cidade Digital---------//
			//---------Rotas de Acompanhamento da Cidade Digital---------//
			//---------Rotas de Ponto da Cidade Digital---------//
			//---------Rotas de Pagamento da Cidade Digital---------//

		//**********Rotas em ClasseEmpenho
		//Rotas de ClaseEmpenho
		//lista classe empenho
		r.HandleFunc().Methods(http.MethodGet)

		//salva classe empenho
		r.HandleFunc().Methods(http.MethodPost)

		//edita classe empenho
		r.HandleFunc().Methods(http.MethodPut)

		//lista classe empenho por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga classe empenho ID
		r.HandleFunc().Methods(http.MethodDelete)

			//----------Rotas De Telefone
				//ROTAS EM TEL_PATH
					//lista telefone
						r.HandleFunc().Methods(http.MethodGet)

					//salva telefone
						r.HandleFunc().Methods(http.MethodPost)
				ROTAS EM TEL_ID
					//apaga telefone ID
						r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Empenho
			//---------Rotas de Empenho
				//ROTAS EM EMPENHO_PATH
					//lista empenho
						r.HandleFunc().Methods(http.MethodGet)

					//salva empenho
						r.HandleFunc().Methods(http.MethodPost)

					//edita empenho
						r.HandleFunc().Methods(http.MethodPut)
				//ROTAS EM EMPENHO_ID_PATH
					//lista empenho por ID
						r.HandleFunc().Methods(http.MethodGet)

					//apaga emepnho ID
						r.HandleFunc().Methods(http.MethodDelete)

			//----------Rotas de Empenho Itens----------//
				//ROTAS EM EMPENHO_ITENS
					//edita empenho itens
						r.HandleFunc().Methods(http.MethodPut)
				//ROTAS EM EMPENHO_ITENS_ID
					//lista empenho itens ID
						r.HandleFunc().Methods(http.MethodGet)


		//**********Rotas em Etapa
		//----------Rotas de Etapa
		//lista etapa
		r.HandleFunc().Methods(http.MethodGet)

		//salva etapa
		r.HandleFunc().Methods(http.MethodPost)

		//edita etapa
		r.HandleFunc().Methods(http.MethodPut)

		//lista etapa por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga etapa ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Fatura
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//

		//**********Rotas em Itens
		//----------Rotas de Itens
		//lista itens
		r.HandleFunc().Methods(http.MethodGet)

		//salva itens
		r.HandleFunc().Methods(http.MethodPost)

		//edita itens
		r.HandleFunc().Methods(http.MethodPut)

		//lista itens por ID (cod_item, cod_tipo_item)
		r.HandleFunc().Methods(http.MethodGet)

		//apaga itens (cod_item, cod_tipo_item)
		r.HandleFunc().Methods(http.MethodDelete)





		//**********Rotas em Modulos
		//----------Rotas de Modulos
		//lista modulos
		r.HandleFunc().Methods(http.MethodGet)

		//lista Modulos Usuario (cod_usuario)
		r.HandleFunc().Methods(http.MethodGet)

		//**********Rotas em Municipios
		//----------Rotas de Municipio
		//lista minicipio
		r.HandleFunc().Methods(http.MethodGet)

		//salva municipio
		r.HandleFunc().Methods(http.MethodPost)

		//edita municipio
		r.HandleFunc().Methods(http.MethodPut)

		//lista municipio por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga municipio ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em NaturezaDespesa
		//----------Rotas de NaturezaDespesa
		//lista natureza despesa
		r.HandleFunc().Methods(http.MethodGet)

		//salva natureza despesa
		r.HandleFunc().Methods(http.MethodPost)

		//edita natureza despesa
		r.HandleFunc().Methods(http.MethodPut)

		//lista natureza despesa por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga natureza despesa ID
		r.HandleFunc().Methods(http.MethodDelete)

		//**********Rotas em Pagamento (otb)
			//----------Rotas de Pagamento
				//ROTAS EM PAG_PATH
					//lista pagamento
						r.HandleFunc().Methods(http.MethodGet)

					//salva pagamento
						r.HandleFunc().Methods(http.MethodPost)

					//edita pagamento
						r.HandleFunc().Methods(http.MethodPut)
				//ROTAS EM PAG_ID_PATH
					//lista pagamento por ID
						r.HandleFunc().Methods(http.MethodGet)

					//apaga pagamento ID
						r.HandleFunc().Methods(http.MethodDelete)

			//----------Rotas de Municipio Fatura
				//ROTAS EM PAG_FAT_MUNICIPIO
					//lista municipio fatura (cd_municipio_cod_ibge)
						r.HandleFunc().Methods(http.MethodGet)
				//ROTAS EM PAG_FAT_LIST
					//lista pagamento fatura (cod_otb)
						r.HandleFunc().Methods(http.MethodGet)
				//ROTAS EM PAG_FAT_SAVE
					//salva pagamento fatura
						r.HandleFunc().Methods(http.MethodPost)

			//----------Rotas de Pagamento Itens
				//ROTAS EM PAG_LIST_ITENS
					//lista pagamento itens	(cod_otb)
						r.HandleFunc().Methods(http.MethodGet)
				//ROTAS EM PAG_EDIT_ITENS
					//edita pagamento itens
						r.HandleFunc().Methods(http.MethodPut)

		//**********Rotas em Prefeitos
		//----------Rotas de Prefeitos
		//lista prefeitos
		r.HandleFunc().Methods(http.MethodGet)

		//salva prefeitos
		r.HandleFunc().Methods(http.MethodPost)

		//edita prefeitos
		r.HandleFunc().Methods(http.MethodPut)

		//lista prefeitos por ID
		r.HandleFunc().Methods(http.MethodGet)

		//apaga prefeitos ID
		r.HandleFunc().Methods(http.MethodDelete)





		//**********Rotas em TipoItem

		//**********Rotas em Tipologia



	*/
	return
}

//Methods: OPTIONS, GET, POST, PUT, DELETE
