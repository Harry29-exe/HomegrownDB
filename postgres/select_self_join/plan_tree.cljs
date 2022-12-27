{PLANNEDSTMT
 :commandType 1
 :queryId 0
 :hasReturning false
 :hasModifyingCTE false
 :canSetTag true
 :transientPlan false
 :dependsOnRole false
 :parallelModeNeeded false
 :jitFlags 0
 :planTree
 {HASHJOIN
  :startup_cost 11.57
  :total_cost 24.20
  :plan_rows 70
  :plan_width 1032
  :parallel_aware false
  :parallel_safe true
  :plan_node_id 0
  :targetlist (
                {TARGETENTRY
                 :expr
                 {VAR
                  :varno 65001
                  :varattno 2
                  :vartype 1043
                  :vartypmod 259
                  :varcollid 100
                  :varlevelsup 0
                  :varnosyn 1
                  :varattnosyn 2
                  :location 7
                  }
                 :resno 1
                 :resname specie
                 :ressortgroupref 0
                 :resorigtbl 16386
                 :resorigcol 2
                 :resjunk false
                 }
                {TARGETENTRY
                 :expr
                 {VAR
                  :varno 65000
                  :varattno 1
                  :vartype 1043
                  :vartypmod 259
                  :varcollid 100
                  :varlevelsup 0
                  :varnosyn 2
                  :varattnosyn 4
                  :location 18
                  }
                 :resno 2
                 :resname surname
                 :ressortgroupref 0
                 :resorigtbl 16392
                 :resorigcol 4
                 :resjunk false
                 }
                )
  :qual <>
  :lefttree
  {SEQSCAN
   :startup_cost 0.00
   :total_cost 11.40
   :plan_rows 140
   :plan_width 516
   :parallel_aware false
   :parallel_safe true
   :plan_node_id 1
   :targetlist (
                 {TARGETENTRY
                  :expr
                  {VAR
                   :varno 1
                   :varattno 1
                   :vartype 20
                   :vartypmod -1
                   :varcollid 0
                   :varlevelsup 0
                   :varnosyn 1
                   :varattnosyn 1
                   :location -1
                   }
                  :resno 1
                  :resname <>
                  :ressortgroupref 0
                  :resorigtbl 0
                  :resorigcol 0
                  :resjunk false
                  }
                 {TARGETENTRY
                  :expr
                  {VAR
                   :varno 1
                   :varattno 2
                   :vartype 1043
                   :vartypmod 259
                   :varcollid 100
                   :varlevelsup 0
                   :varnosyn 1
                   :varattnosyn 2
                   :location -1
                   }
                  :resno 2
                  :resname <>
                  :ressortgroupref 0
                  :resorigtbl 0
                  :resorigcol 0
                  :resjunk false
                  }
                 {TARGETENTRY
                  :expr
                  {VAR
                   :varno 1
                   :varattno 3
                   :vartype 2950
                   :vartypmod -1
                   :varcollid 0
                   :varlevelsup 0
                   :varnosyn 1
                   :varattnosyn 3
                   :location -1
                   }
                  :resno 3
                  :resname <>
                  :ressortgroupref 0
                  :resorigtbl 0
                  :resorigcol 0
                  :resjunk false
                  }
                 )
   :qual <>
   :lefttree <>
   :righttree <>
   :initPlan <>
   :extParam (b)
   :allParam (b)
   :scanrelid 1
   }
  :righttree
  {HASH
   :startup_cost 10.70
   :total_cost 10.70
   :plan_rows 70
   :plan_width 1032
   :parallel_aware false
   :parallel_safe true
   :plan_node_id 2
   :targetlist (
                 {TARGETENTRY
                  :expr
                  {VAR
                   :varno 65001
                   :varattno 1
                   :vartype 1043
                   :vartypmod 259
                   :varcollid 100
                   :varlevelsup 0
                   :varnosyn 2
                   :varattnosyn 4
                   :location -1
                   }
                  :resno 1
                  :resname <>
                  :ressortgroupref 0
                  :resorigtbl 0
                  :resorigcol 0
                  :resjunk false
                  }
                 {TARGETENTRY
                  :expr
                  {VAR
                   :varno 65001
                   :varattno 2
                   :vartype 1043
                   :vartypmod 259
                   :varcollid 100
                   :varlevelsup 0
                   :varnosyn 2
                   :varattnosyn 3
                   :location -1
                   }
                  :resno 2
                  :resname <>
                  :ressortgroupref 0
                  :resorigtbl 0
                  :resorigcol 0
                  :resjunk false
                  }
                 )
   :qual <>
   :lefttree
   {SEQSCAN
    :startup_cost 0.00
    :total_cost 10.70
    :plan_rows 70
    :plan_width 1032
    :parallel_aware false
    :parallel_safe true
    :plan_node_id 3
    :targetlist (
                  {TARGETENTRY
                   :expr
                   {VAR
                    :varno 2
                    :varattno 4
                    :vartype 1043
                    :vartypmod 259
                    :varcollid 100
                    :varlevelsup 0
                    :varnosyn 2
                    :varattnosyn 4
                    :location 18
                    }
                   :resno 1
                   :resname <>
                   :ressortgroupref 0
                   :resorigtbl 0
                   :resorigcol 0
                   :resjunk false
                   }
                  {TARGETENTRY
                   :expr
                   {VAR
                    :varno 2
                    :varattno 3
                    :vartype 1043
                    :vartypmod 259
                    :varcollid 100
                    :varlevelsup 0
                    :varnosyn 2
                    :varattnosyn 3
                    :location 58
                    }
                   :resno 2
                   :resname <>
                   :ressortgroupref 0
                   :resorigtbl 0
                   :resorigcol 0
                   :resjunk false
                   }
                  )
    :qual <>
    :lefttree <>
    :righttree <>
    :initPlan <>
    :extParam (b)
    :allParam (b)
    :scanrelid 2
    }
   :righttree <>
   :initPlan <>
   :extParam (b)
   :allParam (b)
   :hashkeys (
               {RELABELTYPE
                :arg
                {VAR
                 :varno 65001
                 :varattno 2
                 :vartype 1043
                 :vartypmod 259
                 :varcollid 100
                 :varlevelsup 0
                 :varnosyn 2
                 :varattnosyn 3
                 :location 58
                 }
                :resulttype 25
                :resulttypmod -1
                :resultcollid 100
                :relabelformat 2
                :location -1
                }
               )
   :skewTable 16386
   :skewColumn 2
   :skewInherit false
   :rows_total 0
   }
  :initPlan <>
  :extParam (b)
  :allParam (b)
  :jointype 0
  :inner_unique false
  :joinqual <>
  :hashclauses (
                 {OPEXPR
                  :opno 98
                  :opfuncid 67
                  :opresulttype 16
                  :opretset false
                  :opcollid 0
                  :inputcollid 100
                  :args (
                          {RELABELTYPE
                           :arg
                           {VAR
                            :varno 65001
                            :varattno 2
                            :vartype 1043
                            :vartypmod 259
                            :varcollid 100
                            :varlevelsup 0
                            :varnosyn 1
                            :varattnosyn 2
                            :location 67
                            }
                           :resulttype 25
                           :resulttypmod -1
                           :resultcollid 100
                           :relabelformat 2
                           :location -1
                           }
                          {RELABELTYPE
                           :arg
                           {VAR
                            :varno 65000
                            :varattno 2
                            :vartype 1043
                            :vartypmod 259
                            :varcollid 100
                            :varlevelsup 0
                            :varnosyn 2
                            :varattnosyn 3
                            :location 58
                            }
                           :resulttype 25
                           :resulttypmod -1
                           :resultcollid 100
                           :relabelformat 2
                           :location -1
                           }
                          )
                  :location -1
                  }
                 )
  :hashoperators (o 98)
  :hashcollations (o 100)
  :hashkeys (
              {RELABELTYPE
               :arg
               {VAR
                :varno 65001
                :varattno 2
                :vartype 1043
                :vartypmod 259
                :varcollid 100
                :varlevelsup 0
                :varnosyn 1
                :varattnosyn 2
                :location 67
                }
               :resulttype 25
               :resulttypmod -1
               :resultcollid 100
               :relabelformat 2
               :location -1
               }
              )
  }
 :rtable (
           {RTE
            :alias
            {ALIAS
             :aliasname b1
             :colnames <>
             }
            :eref
            {ALIAS
             :aliasname b1
             :colnames ("id" "specie" "uuid")
             }
            :rtekind 0
            :relid 16386
            :relkind r
            :rellockmode 1
            :tablesample <>
            :lateral false
            :inh false
            :inFromCl true
            :requiredPerms 2
            :checkAsUser 0
            :selectedCols (b 9)
            :insertedCols (b)
            :updatedCols (b)
            :extraUpdatedCols (b)
            :securityQuals <>
            }
           {RTE
            :alias
            {ALIAS
             :aliasname u
             :colnames <>
             }
            :eref
            {ALIAS
             :aliasname u
             :colnames ("id" "age" "name" "surname")
             }
            :rtekind 0
            :relid 16392
            :relkind r
            :rellockmode 1
            :tablesample <>
            :lateral false
            :inh false
            :inFromCl true
            :requiredPerms 2
            :checkAsUser 0
            :selectedCols (b 10 11)
            :insertedCols (b)
            :updatedCols (b)
            :extraUpdatedCols (b)
            :securityQuals <>
            }
           {RTE
            :alias <>
            :eref
            {ALIAS
             :aliasname unnamed_join
             :colnames ("id" "specie" "uuid" "id" "age" "name" "surname")
             }
            :rtekind 2
            :jointype 0
            :joinmergedcols 0
            :joinaliasvars <>
            :joinleftcols <>
            :joinrightcols <>
            :lateral false
            :inh false
            :inFromCl true
            :requiredPerms 0
            :checkAsUser 0
            :selectedCols (b)
            :insertedCols (b)
            :updatedCols (b)
            :extraUpdatedCols (b)
            :securityQuals <>
            }
           )
 :resultRelations <>
 :rootResultRelations <>
 :appendRelations <>
 :subplans <>
 :rewindPlanIDs (b)
 :rowMarks <>
 :relationOids (o 16386 16392)
 :invalItems <>
 :paramExecTypes <>
 :utilityStmt <>
 :stmt_location 0
 :stmt_len 0
 }
